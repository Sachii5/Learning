<?php
namespace Data\Company;

class Company {
    public string $name;
}

class Manager extends Company 
{
    function sayHello(string $name)
    {
        echo "Hello, my name is Manager $name" . PHP_EOL;
    }
}

class VicePresident extends Manager 
{
function sayHello(string $name)
    {
        echo "Hello, my name is Vice President $name" . PHP_EOL;
    }

    function sayHelloParent(string $name)
    {
        return parent::sayHello($name);
    }
}
