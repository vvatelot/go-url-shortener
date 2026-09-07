document.addEventListener('DOMContentLoaded', () => {
    setTimeout(function () {
        document.querySelectorAll('.notification').forEach(function ($notification) {
            $notification.classList.add('is-hidden');
        });
    }, 5000);

    const deleteNotification = document.querySelector('.notification .delete')
    if (deleteNotification) {
        deleteNotification.addEventListener('click', () => {
            deleteNotification.parentNode.classList.add('is-hidden');
        });
    }

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
        if (!notification) {
            return;
        }

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

function updateLink(id) {
    var title = document.getElementById("title").value;
    var url = document.getElementById("url").value;

    if (title && url) {
        fetch("/api/links/" + id, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                title: title,
                url: url,
            })
        }).then(function (response) {
            if (!response.ok) {
                return response.text().then(function (message) {
                    throw new Error(message || "Erreur lors de la modification");
                });
            }
            window.location.href = "/?flash=success&message=Le lien a bien été modifié";
        }).catch(function (error) {
            window.location.href = "/?flash=error&message=" + encodeURIComponent(error.message);
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

var clicksChart = null;

function setWindowButtons(window) {
    var button7d = document.getElementById("window-7d");
    var button30d = document.getElementById("window-30d");
    if (!button7d || !button30d) {
        return;
    }
    button7d.classList.toggle("is-primary", window === "7d");
    button30d.classList.toggle("is-primary", window === "30d");
}

function loadClickStats(linkId, window) {
    var canvas = document.getElementById("clicks-chart");
    if (!canvas || typeof Chart === "undefined") {
        return;
    }

    setWindowButtons(window);

    fetch("/api/links/" + linkId + "/clicks/stats?window=" + encodeURIComponent(window))
        .then(function (response) {
            if (!response.ok) {
                throw new Error("Impossible de charger les statistiques");
            }
            return response.json();
        })
        .then(function (data) {
            var totalEl = document.getElementById("stats-total");
            if (totalEl) {
                totalEl.textContent = data.total;
            }

            var labels = data.buckets.map(function (bucket) {
                return bucket.date;
            });
            var values = data.buckets.map(function (bucket) {
                return bucket.count;
            });

            if (clicksChart) {
                clicksChart.destroy();
            }

            clicksChart = new Chart(canvas, {
                type: "bar",
                data: {
                    labels: labels,
                    datasets: [{
                        label: "Clics",
                        data: values,
                        backgroundColor: "rgba(72, 95, 199, 0.65)",
                        borderColor: "rgba(72, 95, 199, 1)",
                        borderWidth: 1
                    }]
                },
                options: {
                    responsive: true,
                    scales: {
                        y: {
                            beginAtZero: true,
                            ticks: {
                                precision: 0
                            }
                        }
                    },
                    plugins: {
                        legend: {
                            display: false
                        }
                    }
                }
            });
        })
        .catch(function (error) {
            console.log(error);
        });
}
