#include <stdio.h>
#include <stdlib.h>
// -------------------------------------------------------------
typedef struct _MemberInfo {
 const char * tag;
 void * addr;
} MemberInfo;
typedef struct _TypeInfo {
  MemberInfo* members;
} TypeInfo;
typedef struct _InterfaceInfo {
	const char** tags;
} InterfaceInfo;
// -------------------------------------------------------------
typedef struct _IReadWriterTbl {
	InterfaceInfo* inter;
	TypeInfo* type;
	int(*Read)(void* this, char* buf, int cb);
	int(*Write)(void* this, char* buf, int cb);
} IReadWriterTbl;
typedef struct _IReadWriter {
	IReadWriterTbl* tab;
	void* data;
} IReadWriter;
InterfaceInfo g_InterfaceInfo_IReadWriter = {
	// ...
};
// -------------------------------------------------------------
typedef struct _A {
	int a;
} A;
int A_Read(A* this, char* buf, int cb) {
	printf("A_Read: %d\n", this->a);
	return cb;
}
int A_Write(A* this, char* buf, int cb) {
	printf("A_Write: %d\n", this->a);
	return cb;
}
TypeInfo g_TypeInfo_A = {
	// ...
};
A* NewA(int params) {
	printf("NewA: %d\n", params);
	A* this = (A*)malloc(sizeof(A));
	this->a = params;
	return this;
}
// -------------------------------------------------------------
typedef struct _B {
	A base;
} B;
int B_Write(B* this, char* buf, int cb) {
	printf("B_Write: %d\n", this->base.a);
	return cb;
}
void B_Foo(B* this) {
	printf("B_Foo: %d\n", this->base.a);
}
TypeInfo g_TypeInfo_B = {
	// ...
};
B* NewB(int params) {
	printf("NewB: %d\n", params);
	B* this = (B*)malloc(sizeof(B));
	this->base.a = params;
	return this;
}
// -------------------------------------------------------------
IReadWriterTbl g_Itbl_IReadWriter_B = {
 &g_InterfaceInfo_IReadWriter,
 &g_TypeInfo_B,
 (int(*)(void* this, char* buf, int cb))A_Read,
 (int(*)(void* this, char* buf, int cb))B_Write
};
int main() {
	B* unnamed = NewB(8);
	IReadWriter p = {
	&g_Itbl_IReadWriter_B,
	unnamed
	};
	p.tab->Read(p.data, NULL, 10);
	p.tab->Write(p.data, NULL, 10);
	return 0;
}
// -------------------------------------------------------------