// package main

// import (
// 	"os"
// 	"strings"
// 	"testing"
// )

// // ---------------------------------------------------------------------------
// // helpers
// // ---------------------------------------------------------------------------

// // buildFakeFile writes content to a temp file and returns its path.
// // The caller is responsible for removing it via t.Cleanup.
// func buildFakeFile(t *testing.T, content string) string {
// 	t.Helper()
// 	f, err := os.CreateTemp("", "banner_test_*.txt")
// 	if err != nil {
// 		t.Fatalf("could not create temp file: %v", err)
// 	}
// 	if _, err := f.WriteString(content); err != nil {
// 		t.Fatalf("could not write temp file: %v", err)
// 	}
// 	f.Close()
// 	t.Cleanup(func() { os.Remove(f.Name()) })
// 	return f.Name()
// }

// // buildValidBannerContent constructs a syntactically valid banner file
// // with the requested number of characters (starting from ASCII 32).
// //
// // Real banner file layout:
// //
// //	""        ← leading separator (index 0) — the file starts with a blank line
// //	<8 lines> ← space char (ASCII 32)
// //	""        ← separator
// //	<8 lines> ← '!' (ASCII 33)
// //	""        ← separator
// //	...
// //
// // Each group of 9 elements is: separator + 8 content lines.
// // The parser skips index 0 of each group (the separator) and reads [1:9].
// func buildValidBannerContent(charCount int) string {
// 	var sb strings.Builder
// 	for i := 0; i < charCount; i++ {
// 		// separator BEFORE each character's 8 lines (matches real file format)
// 		sb.WriteByte('\n')
// 		ch := rune(32 + i)
// 		for line := 0; line < 8; line++ {
// 			sb.WriteString(strings.Repeat(string(ch), 3))
// 			sb.WriteByte('\n')
// 		}
// 	}
// 	return sb.String()
// }

// // ---------------------------------------------------------------------------
// // 1. File-level error cases
// // ---------------------------------------------------------------------------

// func TestLoadBanner_FileNotFound(t *testing.T) {
// 	_, err := LoadBanner("notfound.txt")
// 	if err == nil {
// 		t.Fatal("expected non-nil error for missing file, got nil")
// 	}
// }

// func TestLoadBanner_EmptyFilename(t *testing.T) {
// 	_, err := LoadBanner("")
// 	if err == nil {
// 		t.Fatal("expected error for empty filename, got nil")
// 	}
// }

// func TestLoadBanner_DirectoryInsteadOfFile(t *testing.T) {
// 	dir := t.TempDir()
// 	_, err := LoadBanner(dir)
// 	if err == nil {
// 		t.Fatalf("expected error when path is a directory, got nil")
// 	}
// }

// // ---------------------------------------------------------------------------
// // 2. Content-level error / malformed-file cases
// // ---------------------------------------------------------------------------

// func TestLoadBanner_EmptyFile(t *testing.T) {
// 	path := buildFakeFile(t, "")
// 	_, err := LoadBanner(path)
// 	if err == nil {
// 		t.Fatal("expected error for empty file, got nil")
// 	}
// }

// func TestLoadBanner_TooFewLines_OneCharShort(t *testing.T) {
// 	// Only 7 lines for a single character — incomplete block.
// 	content := strings.Repeat("X\n", 7)
// 	path := buildFakeFile(t, content)
// 	_, err := LoadBanner(path)
// 	if err == nil {
// 		t.Fatal("expected error when block has < 8 lines, got nil")
// 	}
// }

// func TestLoadBanner_MissingSeparatorLine(t *testing.T) {
// 	// Two characters but no blank line between them → 16 lines total, no separator.
// 	// Parser expects groups of 9; this is ambiguous / malformed.
// 	var sb strings.Builder
// 	for i := 0; i < 16; i++ {
// 		sb.WriteString("line\n")
// 	}
// 	path := buildFakeFile(t, sb.String())
// 	_, err := LoadBanner(path)
// 	if err == nil {
// 		t.Fatal("expected error when separator line is missing, got nil")
// 	}
// }

// func TestLoadBanner_Exactly95Characters(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if len(result) != 95 {
// 		t.Fatalf("expected 95 entries in map, got %d", len(result))
// 	}
// }

// func TestLoadBanner_MoreThan95Characters_RejectsOrIgnores(t *testing.T) {
// 	// File has 100 character blocks — implementation must not panic.
// 	// We accept either: error returned OR map has exactly 95 entries.
// 	content := buildValidBannerContent(100)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		return // acceptable — reject outright
// 	}
// 	if len(result) > 95 {
// 		t.Fatalf("map must not exceed 95 printable ASCII characters, got %d", len(result))
// 	}
// }

// func TestLoadBanner_FewerThan95Characters(t *testing.T) {
// 	// Only 10 character blocks — file is incomplete.
// 	content := buildValidBannerContent(10)
// 	path := buildFakeFile(t, content)
// 	_, err := LoadBanner(path)
// 	if err == nil {
// 		t.Fatal("expected error when file contains fewer than 95 characters, got nil")
// 	}
// }

// // ---------------------------------------------------------------------------
// // 3. Map key correctness — ASCII range 32–126
// // ---------------------------------------------------------------------------

// func TestLoadBanner_KeysSpan_ASCII32_To_126(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	for code := 32; code <= 126; code++ {
// 		ch := rune(code)
// 		if _, ok := result[ch]; !ok {
// 			t.Errorf("map missing key for ASCII %d (%q)", code, ch)
// 		}
// 	}
// }

// func TestLoadBanner_NoKeysOutside_ASCII32_To_126(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	for ch := range result {
// 		if ch < 32 || ch > 126 {
// 			t.Errorf("unexpected key in map: ASCII %d (%q)", ch, ch)
// 		}
// 	}
// }

// // ---------------------------------------------------------------------------
// // 4. Per-character slice correctness
// // ---------------------------------------------------------------------------

// func TestLoadBanner_EachCharHasExactly8Lines(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	for ch, lines := range result {
// 		if len(lines) != 8 {
// 			t.Errorf("character %q has %d lines, want 8", ch, len(lines))
// 		}
// 	}
// }

// func TestLoadBanner_SpaceChar_IsPresent(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if _, ok := result[' ']; !ok {
// 		t.Fatal("map must contain the space character (ASCII 32)")
// 	}
// }

// func TestLoadBanner_SpaceChar_Has8Lines(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	lines := result[' ']
// 	if len(lines) != 8 {
// 		t.Fatalf("space character has %d lines, want 8", len(lines))
// 	}
// }

// func TestLoadBanner_TildeChar_IsLastEntry(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if _, ok := result['~']; !ok {
// 		t.Fatal("map must contain '~' (ASCII 126), the last printable character")
// 	}
// }

// func TestLoadBanner_LineContentIsPreserved(t *testing.T) {
// 	// Verify that the exact content of each line is preserved without any
// 	// trimming or modification. We build a file where the space char's 8 lines
// 	// all contain a known string, then assert they come back unchanged.
// 	//
// 	// Format: separator THEN 8 content lines, repeated for all 95 characters.
// 	var sb strings.Builder

// 	// Space char (ASCII 32) — index 0
// 	sb.WriteByte('\n') // leading separator
// 	for line := 0; line < 8; line++ {
// 		sb.WriteString("hello\n")
// 	}

// 	// Pad the remaining 94 characters so the file is valid.
// 	for i := 1; i < 95; i++ {
// 		sb.WriteByte('\n') // separator before each character
// 		for line := 0; line < 8; line++ {
// 			sb.WriteString("x\n")
// 		}
// 	}

// 	path := buildFakeFile(t, sb.String())
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	for i, line := range result[' '] {
// 		if line != "hello" {
// 			t.Errorf("space char line %d: got %q, want %q", i, line, "hello")
// 		}
// 	}
// }

// func TestLoadBanner_SeparatorLineIsNotIncludedInSlice(t *testing.T) {
// 	// Every character slice must have exactly 8 elements, never 9.
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	result, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	for ch, lines := range result {
// 		if len(lines) == 9 {
// 			t.Errorf("character %q slice has 9 entries — separator was not stripped", ch)
// 		}
// 	}
// }

// // ---------------------------------------------------------------------------
// // 5. Real banner files (integration, skipped if files not present)
// // ---------------------------------------------------------------------------

// func testRealBannerFile(t *testing.T, filename string) {
// 	t.Helper()
// 	if _, err := os.Stat(filename); os.IsNotExist(err) {
// 		t.Skipf("banner file %q not found, skipping integration test", filename)
// 	}

// 	result, err := LoadBanner(filename)
// 	if err != nil {
// 		t.Fatalf("LoadBanner(%q) returned unexpected error: %v", filename, err)
// 	}
// 	if len(result) != 95 {
// 		t.Errorf("expected 95 characters, got %d", len(result))
// 	}
// 	for code := 32; code <= 126; code++ {
// 		ch := rune(code)
// 		lines, ok := result[ch]
// 		if !ok {
// 			t.Errorf("missing character ASCII %d (%q)", code, ch)
// 			continue
// 		}
// 		if len(lines) != 8 {
// 			t.Errorf("character ASCII %d (%q): %d lines, want 8", code, ch, len(lines))
// 		}
// 	}
// }

// func TestLoadBanner_StandardTxt(t *testing.T) {
// 	testRealBannerFile(t, "standard.txt")
// }

// func TestLoadBanner_ShadowTxt(t *testing.T) {
// 	testRealBannerFile(t, "shadow.txt")
// }

// func TestLoadBanner_ThinkertoyTxt(t *testing.T) {
// 	testRealBannerFile(t, "thinkertoy.txt")
// }

// // ---------------------------------------------------------------------------
// // 6. Return-value contract
// // ---------------------------------------------------------------------------

// func TestLoadBanner_OnSuccess_ErrorIsNil(t *testing.T) {
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)
// 	_, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("expected nil error on success, got: %v", err)
// 	}
// }

// func TestLoadBanner_OnError_MapIsNil(t *testing.T) {
// 	result, _ := LoadBanner("notfound.txt")
// 	if result != nil {
// 		t.Fatal("expected nil map on error, got non-nil")
// 	}
// }

// func TestLoadBanner_IsDeterministic(t *testing.T) {
// 	// Calling LoadBanner twice on the same file must produce identical maps.
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)

// 	r1, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("first call failed: %v", err)
// 	}
// 	r2, err := LoadBanner(path)
// 	if err != nil {
// 		t.Fatalf("second call failed: %v", err)
// 	}

// 	for ch, lines1 := range r1 {
// 		lines2, ok := r2[ch]
// 		if !ok {
// 			t.Errorf("second call missing key %q", ch)
// 			continue
// 		}
// 		for i, l := range lines1 {
// 			if l != lines2[i] {
// 				t.Errorf("character %q line %d differs between calls: %q vs %q", ch, i, l, lines2[i])
// 			}
// 		}
// 	}
// }

// func TestLoadBanner_ReturnedMap_IsNotShared(t *testing.T) {
// 	// Mutating the returned map must not affect a second call.
// 	content := buildValidBannerContent(95)
// 	path := buildFakeFile(t, content)

// 	r1, _ := LoadBanner(path)
// 	r1['A'][0] = "MUTATED"

//		r2, _ := LoadBanner(path)
//		if r2['A'][0] == "MUTATED" {
//			t.Fatal("LoadBanner returned a shared mutable reference; maps must be independent")
//		}
//	}
package main
