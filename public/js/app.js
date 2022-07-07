function switchLink(id) {
    if (confirm("Êtes vous sûr de vouloir changer l'état de ce lien ?")) {
        fetch("/api/links/" + id + "/switch-active", {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
        });
    }
}

function addNewLink(id) {
    var url = document.getElementById("url").value;

    if (url) {
        fetch("/api/links", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                url: url,
            })
        }).then(function (response) {
            window.location.href = "/";
        }).catch(function (error) {
            console.log(error);
        })
    }
}

function deleteLink(id) {
    if (confirm("Êtes vous sûr de vouloir supprimer ce lien ?")) {
        fetch("/api/links/" + id, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
        }).then(function (response) {
            location.reload();
        }).catch(function (error) {
            console.log(error);
        })
    }
}