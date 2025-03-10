document.addEventListener('DOMContentLoaded', function () {
    const newEmailRow = document.getElementById('newEmailRow');
    const newEmailButton = document.getElementById('newEmailButton');
    newEmailButton.addEventListener('click', function () {
        newEmailRow.classList.add('show');
    });
    const closeEmailButton = document.getElementById('closeNewEmail');
    closeEmailButton.addEventListener('click', function () {
        newEmailRow.classList.remove('show');
    });
    // when 'send' is pressed
    const form = document.getElementById('newEmailForm');
    form.addEventListener('submit', function (event) {
        event.preventDefault();
        // get the form data
        const formData = new FormData(form);
        const id_element = document.getElementById('case_id');
        var case_id = id_element.value ? id_element.value : '';
        // create email instance
        const newEmail = {
            Case_id: case_id,
            From: formData.get('from'),
            To: formData.get('to'),
            Cc: formData.get('cc'),
            Subject: formData.get('subject'),
            Content: formData.get('message'),
            Actioned: false
        };
        // post email to api endpoint '/api/add-email'
        fetch('/api/add-email', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newEmail)
        })
            .then(response => response.json())
            .then(data => {
            console.log('Success:', data);
        })
            .catch((error) => {
            console.error('Error:', error);
        });
        // send email
        window.location.href = '/case.html?case_id=' + case_id;
    });
});
export {};
