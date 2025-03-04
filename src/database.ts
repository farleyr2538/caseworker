export interface email {
    Id? : number;
    Case_id : number;
    Datetime : Date;
    From? : string;
    To : string;
    Cc? : string;
    Subject : string;
    Content : string;
    Actioned : boolean;
}

export interface constituent {
    id : number;
    first_name : string;
    last_name : string;
    email : string;
    phone : string;
    address1 : string;
    address2 : string;
    area : string;
    city : string;
    postcode : string;
} 

export interface case_ {
    id : number;
    constituent_id : number;
    reference : string | undefined;
    category : string;
    summary : string;
}

export const data = {
    emails : [
        {
            Id: 1,
            Case_id: 1,
            Datetime: new Date(),
            From: 'john_doe@gmail.com',
            To: 'jane_doe@yahoo.com',
            Cc: '',
            Subject: 'Test Email',
            Content: 'This is a test email',
            Actioned: false
        }
    ],
    constituents : [
        {
            id: 1,
            first_name: 'John',
            last_name: 'Doe',
            email: 'john_doe@gmail.com',
            phone: '0123456789',
            address1: '1 Test Street',
            address2: '',
            area: 'Test Area',
            city: 'Test City',
            postcode: 'TE5 7PC'
        },
        {
            id: 2,
            first_name: 'Jane',
            last_name: 'Doe',
            email: 'jane_doe@yahoo.com',
            phone: '9876543210',
            address1: '2 Test Street',
            address2: '',
            area: 'Test Area',
            city: 'Test City',
            postcode: 'TE5 7PC'
        }
    ],
    cases : [
        {
            id: 1,
            constituent_id: 1,
            category: 'General',
            summary: 'This is a test case'
        }
    ]
}