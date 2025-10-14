<?php 

require_once "data/Person.php";

use Data\Human\Person;

$person = new Person("Alber", null);
// jadi value $name bukan diakses dari class person tapi dari object $person
$person->name = "Alber";

$person->sayHello("Jihan");
$person->sayHello(null);
$person->info();
?>