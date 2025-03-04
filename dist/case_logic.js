document.addEventListener('DOMContentLoaded', function () {
    const newEmailButton = document.getElementById('newEmailButton');
    newEmailButton.addEventListener('click', function () {
        // show message
        const newEmailRow = document.getElementById('newEmailRow');
        // newEmailRow.classList.remove('d-none');
        newEmailRow.classList.add('show');
    });
    // when 'send' is pressed
    const form = document.getElementById('newEmailForm');
    form.addEventListener('submit', function (event) {
        event.preventDefault();
        // get the form data
        const formData = new FormData(form);
        const id_element = document.getElementById('case_id');
        const case_id = parseInt(id_element.value);
        // create email instance
        const newEmail = {
            Case_id: case_id,
            Datetime: new Date(),
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
