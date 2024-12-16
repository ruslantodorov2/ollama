NVCC?=nvcc

NVCC_PREPEND_FLAGS= \
	-t 2 \
	-split-compile 0 \
	-std=c++17 \
	-I.. \
	-I../include \
	$(foreach ARCH,$(subst ;, ,$(CUDA_ARCHS)),--generate-code=arch=compute_$(ARCH),code=sm_$(ARCH)) \

NVCC_APPEND_FLAGS= \
	-DGGML_USE_CUDA \
	# -DGGML_CUDA_USE_GRAPHS=1 \

OBJECTS+= \
	$(patsubst %.cu,%.o,$(wildcard *.cu)) \
	$(patsubst %.cu,%.o,$(wildcard template-instances/*.cu)) \

libggml-cuda_v11.so: CUDA_ARCHS?=50;52;53;60;61;62;70;72;75;80;86
libggml-cuda_v12.so: CUDA_ARCHS?=60;61;62;70;72;75;80;86;87;89;90;90a

%.o: %.cu
	$(NVCC) $(NVCC_PREPEND_FLAGS) -c $< -o $@ $(NVCC_APPEND_FLAGS)

%.so: $(OBJECTS)
	$(CC) $(LDFLAGS) -shared -o $@ $^

clean:
	$(RM) $(OBJECTS) *.so
