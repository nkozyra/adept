<?php

	$questions = [];
	$questions[] = [ 
		"text" => "Dual Photography Is:",
		"options" => [
			[
				"text"=>"Controlling light sources allowing more data than traditional photography",
				"correct" => true
			],
			[
				"text"=>"The use of multiple cameras to produce a novel image",
				"correct" => false
			],
			[
				"text"=>"Doubling the number of lenses to meet light source",
				"correct" => false
			]						
		]
	];


	$questions[] = [ 
		"text" => "Dual photography mostly came from this university",
		"options" => [
			[
				"text"=>"Harvard",
				"correct" => false
			],
			[
				"text"=>"Georgia Tech",
				"correct" => false
			],
			[
				"text"=>"Stanford",
				"correct" => true
			]						
		]
	];


	$questions[] = [ 
		"text" => "Which surface type is a diffuse surface? ",
		"options" => [
			[
				"text"=>"Specular",
				"correct" => false
			],
			[
				"text"=>"Mirrored",
				"correct" => false
			],
			[
				"text"=>"Matte",
				"correct" => true
			]						
		]
	];


	$questions[] = [ 
		"text" => "What are the elements of computational photography: ",
		"options" => [
			[
				"text"=>"Illumination, Import, Modification Sharing",
				"correct" => false
			],
			[
				"text"=>"Illumination, Optics, Sensor, Processing, Display",
				"correct" => true
			],
			[
				"text"=>"Capturing, Importing, Editing, Printing",
				"correct" => true
			]						
		]
	];

	$questions[] = [ 
		"text" => "What are the parts of a novel camera:",
		"options" => [
			[
				"text"=>"Generalized optics, Aperture, Sensor, Processing",
				"correct" => true
			],
			[
				"text"=>"Sensor, Lens, Filters, Functions",
				"correct" => false
			],
			[
				"text"=>"Camera, Software",
				"correct" => false
			]						
		]
	];			

	$f = fopen('assets/js/questions.json','w+');
	fwrite($f,json_encode($questions));
	fclose($f);