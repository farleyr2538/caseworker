import { constituent, email, case_, data } from './database';

document.addEventListener('DOMContentLoaded', function() {

    const newEmailRow : any = document.getElementById('newEmailRow');

    const newEmailButton : any = document.getElementById('newEmailButton');
    newEmailButton.addEventListener('click', function() {        
        newEmailRow.classList.add('show');
    })

    const closeEmailButton : any = document.getElementById('closeEmailButton');
    closeEmailButton.addEventListener('click', function() {
        newEmailRow.classList.remove('show');
    })
    
    // when 'send' is pressed
    const form : any = document.getElementById('newEmailForm');
    form.addEventListener('submit', function(event : SubmitEvent) {
        event.preventDefault();

        // get the form data
        const formData  = new FormData(form);
        const id_element : any = document.getElementById('case_id');
        const case_id : number = parseInt(id_element.value);

        // create email instance
        const newEmail: email = {
            Case_id: case_id,
            Datetime: new Date(),
            From: formData.get('from') as string | undefined,
            To: formData.get('to') as string,
            Cc: formData.get('cc') as string | undefined,
            Subject: formData.get('subject') as string,
            Content: formData.get('message') as string,
            Actioned: false
        }

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
})