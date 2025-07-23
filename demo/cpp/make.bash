#/bin/sh
clang $(llvm-config --cxxflags --ldflags --libs) -lclang main.cpp