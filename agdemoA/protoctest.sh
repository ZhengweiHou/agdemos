#!/bin/bash
aggo proto -a -p all -e idl idl



aggo proto -p all -m server -e ./idl/api  ./idl/api/demoa

aggo proto -p kitex,hertz -m client -e ./idl/api  ./idl/api/demob
