var links = document.querySelectorAll("nav a");
for (var i = 0; i < links.length; i++) {
	var link = links[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("active");
		break;
	}
}