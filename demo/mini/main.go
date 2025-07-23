package main

import (
	_ "unsafe"

	"github.com/goplus/lib/c"
)

const (
	LLGoFiles   = "$(llvm-config --cflags): wrap/wrap.cpp"
	LLGoPackage = "link: -L$(llvm-config --libdir) -lclang; -lclang"
)

type Index struct {
	Unused [0]byte
}

type TranslationUnit struct {
	Unused [0]byte
}

type UnsavedFile struct {
	/**
	 * The file whose contents have not yet been saved.
	 *
	 * This file must already exist in the file system.
	 */
	Filename *c.Char

	/**
	 * A buffer containing the unsaved contents of this file.
	 */
	Contents *c.Char

	/**
	 * The length of the unsaved contents of this buffer.
	 */
	Length c.Ulong
}

//go:linkname CreateIndex C.clang_createIndex
func CreateIndex(excludeDeclarationsFromPCH, displayDiagnostics c.Int) *Index

// llgo:link (*Index).ParseTranslationUnit C.clang_parseTranslationUnit
func (*Index) ParseTranslationUnit(
	sourceFilename *c.Char, commandLineArgs **c.Char, numCommandLineArgs c.Int,
	unsavedFiles *UnsavedFile, numUnsavedFiles c.Uint, options c.Uint) *TranslationUnit {
	return nil
}

type Cursor struct {
	Kind  c.Int
	xdata c.Int
	data  [3]c.Pointer
}

//llgo:link (*TranslationUnit).Cursor C.clang_getTranslationUnitCursor
func (t *TranslationUnit) Cursor() (ret Cursor) {
	return
}

type ClientData c.Pointer

//go:linkname VisitChildren C.clang_visitChildren
func VisitChildren(
	root Cursor,
	fn Visitor,
	clientData ClientData) (ret c.Uint) {
	return
}

//go:linkname visitChildrenCallback C.visitChildrenCallback
func visitChildrenCallback(cursor, parent Cursor, clientData ClientData) c.Int

func visit(cursor, parent Cursor, clientData ClientData) c.Int {
	println("visit", cursor.Kind, parent.Kind)
	return 0
}

//llgo:type C
type Visitor func(cursor, parent Cursor, clientData ClientData) c.Int

func main() {
	index := CreateIndex(0, 0)
	sourceFile := *c.Advance(c.Argv, 1)
	tu := index.ParseTranslationUnit(sourceFile, nil, 0, nil, 0, 0)

	root := tu.Cursor()
	println(tu, root.Kind, root.xdata)
	VisitChildren(root, visitChildrenCallback, nil)
}

// func parse(filename *c.Char) {
// 	index := clang.CreateIndex(0, 0)
// 	args := make([]*c.Char, 3)
// 	args[0] = c.Str("-x")
// 	args[1] = c.Str("c++")
// 	args[2] = c.Str("-std=c++11")
// 	unit := index.ParseTranslationUnit(
// 		filename,
// 		unsafe.SliceData(args), 3,
// 		nil, 0,
// 		clang.DetailedPreprocessingRecord,
// 	)

// 	if unit == nil {
// 		println("Unable to parse translation unit. Quitting.")
// 		c.Exit(1)
// 	}

// 	context.setUnit(unit)
// 	cursor := unit.Cursor()

// 	println("======> VisitChildren")
// 	clang.VisitChildren(cursor, visit, nil)
// 	unit.Dispose()
// 	index.Dispose()
// }

// func main() {
// 	if c.Argc != 2 {
// 		fmt.Fprintln(os.Stderr, "Usage: <C++ header file>\n")
// 		return
// 	} else {
// 		sourceFile := *c.Advance(c.Argv, 1)
// 		parse(sourceFile)
// 	}
// }
