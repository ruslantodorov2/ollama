CPPFLAGS+=-DGGML_USE_BLAS \

OBJECTS+=ggml-blas.o

%.so: $(OBJECTS)
	$(CC) $(LDFLAGS) -shared -o $@ $^

clean:
	$(RM) $(OBJECTS) *.so
