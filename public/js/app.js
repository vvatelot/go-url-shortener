document.addEventListener('DOMContentLoaded', () => {
    (document.querySelectorAll('.notification .delete') || []).forEach(($delete) => {
        const $notification = $delete.parentNode;

        $delete.addEventListener('click', () => {
            $notification.parentNode.removeChild($notification);
        });
    });
});

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

function addNewLink() {
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
        }).then(function () {
            window.location.href = "/?flash=success&message=Le lien a bien été ajouté";
        }).catch(function (error) {
            window.location.href = "?flash=error&message=" + error.message;
            console.log(error);
        })
    }

    return false;
}

function deleteLink(id) {
    if (confirm("Êtes vous sûr de vouloir supprimer ce lien ?")) {
        fetch("/api/links/" + id, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
        }).then(function () {
            window.location.href = "/?flash=success&message=Le lien a bien été supprimé";
        }).catch(function (error) {
            console.log(error);
        })
    }
}