# client-go

# k8s object in go code meaning

if any go struct implements runtime.Object interface from apimachinery pkg = becomes k8s object/resource

https://github.com/kubernetes/apimachinery/blob/master/pkg/runtime/interfaces.go

check for type Object u will get 3 methods.

out of 3 method we will discuss about DeepCopyObject() Object

example: check ur lister code and check pod struct.

type TypeMeta struct have 2 methods - u can check by click on "find all implementations" by right click on 2 methods (SetGroupVersionKind, GroupVersionKind)

# Part 4

no watch verb ? - to reduce load on api server (only one resource will make n requests to check events)
instead use informers - use watch internally but efficiently leverages in memory store.
shared informer factory - to reduce the load on api servers (from n informers watching n group versions )
queues: in informers we can register some functions e.g. addfunc, deletefunc, updatefunc
        these functions have business logics like if we create pod p1 then addfunc is going to do lot of things i.e. n things so we enqueue p1 event's addfunc tasks.

# Part 5

API Machinery
=============

GVK: kind like deployment, rs
GVR: kind / http endpoint has resource  -> plural





