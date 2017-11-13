// Anchor links
var contents = ['','#','##','###','####','#####','######']
var anchorForId = function (id,level) {
  var anchor = document.createElement("a");
  anchor.className = "header-link";
  anchor.href      = "#" + id;
  anchor.innerHTML = "<i>" + contents[level]+"</i>";
  return anchor;
};

var linkifyAnchors = function (level) {
  var headers = document.getElementsByTagName("h" + level);
  for (var h = 0; h < headers.length; h++) {
    var header = headers[h];

    if (typeof header.id !== "undefined" && header.id !== "") {
      header.appendChild(anchorForId(header.id,level));
    }
  }
};

document.onreadystatechange = function () {
  if (this.readyState === "complete") {
    for (var level = 2; level <= 6; level++) {
      linkifyAnchors(level);
    }
  }
};