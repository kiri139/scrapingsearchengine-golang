package materials

/*

  #include <dlfcn.h>
  #include <stdio.h>
  #include <stdlib.h>

  char* Cretumecab;
  void CGOnilstringfunc();

  typedef void*(*MECABNEWFUNC)(char*);
  typedef char*(*MECABPARSEFUNC)(void*, char*);

  char* linuxapimecabfunc(char* optin, char* input)
  {

  	void* openfile = dlopen("libmecab.so", RTLD_LAZY);

  	if openfile == NULL
  	{
  	  printf("MECABSOERROR  libmecab.so");
  	}

  	if (openfile)
  	{

      MECABNEWFUNC dllmecabnewfunc = (MECABNEWFUNC)dlsym(OpenDLL, "mecab_new2");
      MECABPARSEFUNC dllmecabparsefunc = (MECABPARSEFUNC)dlsym(OpenDLL, "mecab_sparse_tostr");

      if (dllmecabnewfunc == NULL || dllmecabparsefunc == NULL)
      {
        printf("MECABFUNCERROR");
      }

      if (dllmecabnewfunc && dllmecabparsefunc)
      {
        void *mecaboption = ((void *)(*dllmecabnewfunc)(optin));
        Cretumecab = (char *)(*dllmecabparsefunc)(mecaboption, input);
      }

  	}

  	dlclose(openfile);

  	return Cretumecab

  }

  void CGOnilstringfunc(){
    Cretumecab = "";
  }

*/
import "C"

import (
  "sort"
  "strings"
  "unicode/utf8"
  "unsafe"
)

type (

  Listmectysts struct {
    Names string
    Counts int
  }

  Listmectystszyun []*Listmectysts

)

func (a Listmectystszyun) Len()          int  { return len(a)}
func (a Listmectystszyun) Less(i, j int) bool { return a[i].Counts < a[j].Counts }
func (a Listmectystszyun) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func LinuxHashtags(s string) (resut1,result2,result3 string) {
  var (
  	maplst         = make(map[string]int)
    mapdetermining = make(map[string]bool)
    Listsslissts   = make([]*Listmectysts, 0)
  )

  inputs, optins := C.CString(s), C.CString("-Owakati")
  
  // jp
  wakati := C.GoString(C.mains(inputs, optins))

  C.CGOnilstringfunc()

  CGOdllsplit := strings.Split(wakati, " ")

  for _, slli := range CGOdllsplit {
    maplst[slli]++
  }

  for val, key := range maplst {
    if utf8.RuneCountInString(val) == 1 {
      continue
    }
    if _, ok := mapdetermining[val]; !ok {
      mapdetermining[val] = true
      if a := mecabpp(val); strings.Contains(a, "名詞") == true || strings.Contains(a, "感動詞") == true {
        Listsslissts = append(Listsslissts, &Listmectysts{Names: val,Counts: key})
      }
    }
  }

  C.CGOnilstringfunc()
  C.free(unsafe.Pointer(inputs))
  C.free(unsafe.Pointer(optins))
  sort.Sort(sort.Reverse(Listmectystszyun(Listsslissts)))

  slilen := len(Listsslissts)

  switch slilen {
  case 0:
    return "", "", ""
  case 1:
    return Listsslissts[0].Names, "",""
  case 2:
    return Listsslissts[0].Names, Listsslissts[1].Names, ""
  default:
    return Listsslissts[0].Names, Listsslissts[1].Names, Listsslissts[2].Names
  }
  return
}

func mecabpp(s string)string{
  input, null := C.CString(s), C.CString("NULL")
   
  rs := C.GoString(C.mains(input, null))

  C.free(unsafe.Pointer(input))
  C.free(unsafe.Pointer(null))

  return rs
}
