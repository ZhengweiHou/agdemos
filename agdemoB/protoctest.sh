#!/bin/bash
aggo proto -a -p all -e idl idl



aggo proto -p service,hertz,go,api,server -m server -e ./idl/api  ./idl/api