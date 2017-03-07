package main

import "gopkg.in/kataras/iris.v6"

func HandleRequest(ctx *iris.Context) {
    ctx.Writef(ctx.GetString("path"));
}
