#!/bin/bash

ScriptPath=$(cd `dirname $0` && pwd)
ProjectPath=$ScriptPath/..
GitBlogPath=$ProjectPath/mdgsf.github.io

cd $GitBlogPath && git pull

supervisorctl restart beegoBlog
