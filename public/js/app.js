document.addEventListener('DOMContentLoaded', () => {
    setTimeout(function () {
        document.querySelectorAll('.notification').forEach(function ($notification) {
            $notification.classList.add('is-hidden');
        });
    }, 5000);

    const deleteNotification = document.querySelector('.notification .delete')
    deleteNotification.addEventListener('click', () => {
        deleteNotification.parentNode.classList.add('is-hidden');
    });

    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

    $navbarBurgers.forEach(el => {
        el.addEventListener('click', () => {

            const target = el.dataset.target;
            const $target = document.getElementById(target);

            el.classList.toggle('is-active');
            $target.classList.toggle('is-active');

        });
    });
});


function copyLink(key) {
    const url = window.location.origin + "/r/" + key;
    navigator.clipboard.writeText(url).then(function () {
        const notification = document.querySelector('.notification')

        notification.classList.remove('is-hidden', 'is-danger', 'is-success');
        notification.classList.add('is-primary');
        notification.querySelector('span').innerHTML = "Lien copié dans le presse-papier";
        setTimeout(function () {
            notification.classList.add('is-hidden');
        }, 5000)
    })
}

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