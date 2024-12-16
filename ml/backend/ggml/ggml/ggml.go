package ggml

// #cgo CXXFLAGS: -std=c++17
// #cgo CPPFLAGS: -I${SRCDIR}/include -I${SRCDIR}/ggml-cpu -DNDEBUG
import "C"
