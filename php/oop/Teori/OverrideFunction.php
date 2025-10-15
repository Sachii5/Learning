<?php
require_once "data/Company.php";

use Data\Company\Manager;
use Data\Company\VicePresident;

$manager = new Manager();
$vp = new VicePresident();

$manager->sayHello("Alber");
$vp->sayHello("Wildan");
$vp->sayHelloParent("Alber");