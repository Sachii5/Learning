<?php

class Family
{
    public string $name;

    function __construct(string $name)
    {
        $this->name = $name;
    }
}

class Child extends Family{}
class GrandChild extends Family{}

function sayName(Family $family)
{
    echo "My name is $family->name" . PHP_EOL;
}