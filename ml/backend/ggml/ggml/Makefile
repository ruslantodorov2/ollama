SRCDIR=$(realpath .)

CXXFLAGS= \
	-std=c++11 \

CPPFLAGS= \
	-O3 \
	-I$(SRCDIR) \
	-I$(SRCDIR)/include \
	-I$(SRCDIR)/ggml-cpu \
	-fPIC \
	-DNDEBUG \
	-DGGML_BACKEND_DL \
	-DGGML_BACKEND_SHARED \

LDFLAGS= \
	-lstdc++ \

export CXXFLAGS CPPFLAGS LDFLAGS

OBJECTS= \
	$(SRCDIR)/ggml.o \
	$(SRCDIR)/ggml-alloc.o \
	$(SRCDIR)/ggml-backend.o \
	$(SRCDIR)/ggml-opt.o \
	$(SRCDIR)/ggml-threading.o \
	$(SRCDIR)/ggml-quants.o \

export OBJECTS

ifeq ($(shell uname -p),arm)
CPPFLAGS+= \
	-D__aarch64__ \
	-D__ARM_NEON \
	-D__ARM_FEATURE_FMA \
	-DGGML_USE_CPU_AARCH64 \

OBJECTS+= \
	$(SRCDIR)/ggml-aarch64.o
endif

ifeq ($(shell uname),Darwin)
CPPFLAGS+= \
	-DACCELERATE_NEW_LAPACK \
	-DACCELERATE_LAPACK_ILP64 \
	-DGGML_USE_ACCELERATE \
	-DGGML_BLAS_USE_ACCELERATE \

LDFLAGS+=-framework Accelerate
endif

cpu cpu_avx cpu_avx2: $(OBJECTS)
	$(MAKE) -C ggml-cpu libggml-$@.so

metal: $(OBJECTS)
	$(MAKE) -C ggml-metal libggml-$@.so

cuda_v11 cuda_v12: $(OBJECTS)
	$(MAKE) -C ggml-cuda libggml-$@.so

blas: $(OBJECTS)
	$(MAKE) -C ggml-blas libggml-$@.so

clean:
	$(MAKE) -C ggml-blas $@
	$(MAKE) -C ggml-cpu $@
	$(MAKE) -C ggml-cuda $@
	$(MAKE) -C ggml-metal $@
