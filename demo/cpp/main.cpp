#include <clang-c/Index.h>
#include <stdio.h>

// 定义回调函数
extern "C"
CXChildVisitResult visitChildrenCallback(CXCursor cursor, 
                                       CXCursor parent, 
                                       CXClientData client_data) {
    // 获取光标的位置和名称
    printf("visit\n");
    return CXChildVisit_Recurse;
//    CXString cursorKindName = clang_getCursorKindSpelling(clang_getCursorKind(cursor));
//    CXString cursorSpelling = clang_getCursorSpelling(cursor);
//    CXSourceLocation location = clang_getCursorLocation(cursor);
    
//    CXString filename;
//    unsigned line, column;
//    clang_getPresumedLocation(location, &filename, &line, &column);
    
//    printf("Found %s '%s' at %s:%d:%d\n", 
//           clang_getCString(cursorKindName),
//           clang_getCString(cursorSpelling),
//           clang_getCString(filename),
//           line, column);
    
//    // 释放资源
//    clang_disposeString(cursorKindName);
//    clang_disposeString(cursorSpelling);
//    clang_disposeString(filename);
    
//    // 继续遍历子节点
//    return CXChildVisit_Recurse;
}

int main(int argc, char *argv[]) {
    if (argc < 2) {
        fprintf(stderr, "Usage: %s <filename>\n", argv[0]);
        return 1;
    }
    
    // 创建索引
    CXIndex index = clang_createIndex(0, 0);
    
    // 解析文件
    CXTranslationUnit tu = clang_parseTranslationUnit(
        index,
        argv[1],
        NULL, 0,  // 命令行参数
        NULL, 0,  // 未保存文件
        CXTranslationUnit_None);
    
    if (!tu) {
        fprintf(stderr, "Unable to parse translation unit\n");
        return 1;
    }
    
    // 获取根光标
    CXCursor rootCursor = clang_getTranslationUnitCursor(tu);
    
    // 遍历子节点
    clang_visitChildren(
        rootCursor,
        visitChildrenCallback,
        NULL);  // 客户端数据
    
    // 清理
    clang_disposeTranslationUnit(tu);
    clang_disposeIndex(index);
    
    return 0;
}