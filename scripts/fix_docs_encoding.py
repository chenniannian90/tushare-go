#!/usr/bin/env python3
"""
Fix encoding issues in Markdown documentation files.
This script attempts to recover text corrupted by encoding problems.
"""

import os
import sys
import re
from pathlib import Path

def fix_mojibake(text):
    """
    Attempt to fix mojibake (double-encoded or corrupted text).
    This is a common issue where UTF-8 text was incorrectly decoded as latin-1.
    """
    try:
        # Try common encoding fixes
        # Case 1: UTF-8 decoded as latin-1, then encoded as UTF-8 again
        try:
            fixed = text.encode('latin-1').decode('utf-8')
            if '��' not in fixed and len(fixed) == len(text):
                return fixed
        except (UnicodeDecodeError, UnicodeEncodeError):
            pass

        # Case 2: Try to recover by ignoring/replacing bad characters
        fixed = text.encode('utf-8', errors='replace').decode('utf-8', errors='replace')
        return fixed

    except Exception:
        return text

def clean_replacement_chars(text):
    """Remove Unicode replacement characters and try to recover."""
    # Remove replacement characters
    cleaned = text.replace('', '')

    # Remove other control characters except newlines and tabs
    cleaned = re.sub(r'[\x00-\x08\x0B\x0C\x0E-\x1F\x7F]', '', cleaned)

    return cleaned

def fix_file(filepath):
    """Fix encoding issues in a single file."""
    print(f"Processing: {filepath.name}")

    try:
        # Read file with error handling
        with open(filepath, 'r', encoding='utf-8', errors='replace') as f:
            content = f.read()

        original_replacement_count = content.count('')

        if original_replacement_count == 0:
            print(f"  ✅ No issues found")
            return False

        print(f"  Found {original_replacement_count} replacement characters")

        # Attempt to fix encoding
        fixed_content = fix_mojibake(content)

        # Clean up replacement characters
        fixed_content = clean_replacement_chars(fixed_content)

        # Check if we actually fixed something
        new_replacement_count = fixed_content.count('')

        if new_replacement_count < original_replacement_count:
            # Write fixed content back
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(fixed_content)

            print(f"  ✅ Fixed: {original_replacement_count} → {new_replacement_count} (removed {original_replacement_count - new_replacement_count} bad chars)")
            return True
        else:
            print(f"  ⚠️  Could not fully fix (still has {new_replacement_count} replacement chars)")
            return False

    except Exception as e:
        print(f"  ❌ Error: {e}")
        return False

def main():
    """Main function to process all documentation files."""
    docs_dir = Path("docs")

    if not docs_dir.exists():
        print(f"❌ Error: {docs_dir} directory not found")
        sys.exit(1)

    # Find all markdown files
    md_files = list(docs_dir.glob("*.md"))

    if not md_files:
        print(f"❌ No markdown files found in {docs_dir}")
        sys.exit(1)

    print(f"🔍 Found {len(md_files)} markdown files")
    print()

    total_files = 0
    fixed_files = 0

    for filepath in sorted(md_files):
        total_files += 1
        if fix_file(filepath):
            fixed_files += 1
        print()

    print("=" * 60)
    print(f"📊 Summary:")
    print(f"   Total files: {total_files}")
    print(f"   Files fixed: {fixed_files}")
    print(f"   Success rate: {fixed_files/total_files*100:.1f}%")

    if fixed_files > 0:
        print(f"\n✅ Successfully fixed {fixed_files} file(s)")
    else:
        print(f"\n⚠️  No files were fixed")

if __name__ == "__main__":
    main()