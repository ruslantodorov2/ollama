CPPFLAGS+= \
	-DGGML_USE_METAL \
	-DGGML_METAL_EMBED_LIBRARY \

LDFLAGS+= \
	-framework Foundation \
	-framework Metal \
	-framework MetalKit \
	-framework Accelerate \

OBJECTS+= \
	ggml-metal-embed.o \
	ggml-metal.o \

%.so: $(OBJECTS)
	$(CC) $(LDFLAGS) -shared -o $@ $^

clean:
	$(RM) $(OBJECTS) *.so
