CPPFLAGS+= \
	-Iamx \
	-Illamafile \
	-DGGML_USE_CPU \

SOURCES= \
	$(wildcard *.c) \
	$(wildcard **/*.c) \
	$(wildcard *.cpp) \
	$(wildcard **/*.cpp) \

OBJECTS+=$(addsuffix .o,$(SOURCES))

%.c.o: %.c
	$(CC) $(CPPFLAGS) $(CFLAGS) -c -o $@ $<

%.cpp.o: %.cpp
	$(CXX) $(CPPFLAGS) $(CXXFLAGS) -c -o $@ $<

%.so: $(OBJECTS)
	$(CC) $(LDFLAGS) -shared -o $@ $^

libggml-cpu_avx.so: CPPFLAGS+=-mavx
libggml-cpu_avx2.so: CPPFLAGS+=-mavx2 -mfma -mf16c

clean:
	$(RM) $(OBJECTS) *.so
