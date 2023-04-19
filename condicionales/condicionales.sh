#!/bin/bash
# Author: Ramón Pascual
# Script interactivo - condicionales
name=""
edad=0
year=0
read -r "Ingresa tu nombre: " name
read -r "Ingresa tu edad: " edad
read -r "Ingresa el año actual: " year

echo "Hola, $name, tu edad es $edad años."
echo "===================================="
if (("$edad" >= 18)); then
	echo "Eres mayor de edad"
else
	echo "$name, no eres mayor de edad "
fi

echo "========================================="

if [ "$edad" -ge 18 ] && [ "$year" -eq 2022 ]; then
	echo "$name, puedes votar"
else
	echo "no puedes votar"
fi