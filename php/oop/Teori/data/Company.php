<?php
namespace Data\Company;

class Company {
    public string $name;
}

class Manager extends Company {
    function sayHello(string $name): void
    {
        echo "Hello, my name is Manager $name" . PHP_EOL;
    }
}

class VicePresident extends Company 
{
    function sayHello(string $name): void
    {
        echo "Hello, my name is Vice President $name" . PHP_EOL;
    }
}
