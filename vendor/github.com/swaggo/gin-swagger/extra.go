package ginSwagger

var (
	extraJS = `
	var bg = document.getElementById('dg0');
	console.log("bg:", bg);
	console.log("extra.js => init");

	var bt =document.createElement("button");
	bt.innerHTML = '172.16.101.185';
	bt.onclick = function () {
		console.log('test click!!');
	};
	bg.appendChild(bt);

	`
)
