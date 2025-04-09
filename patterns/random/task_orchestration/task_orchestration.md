## Task Orchestration

Given an API contract like below

POST Get Loan Details /api/v1/loan/details

Request Body
```json
{
    "clientApplicationId": "APP1234",
    "partnershipApplicationId": "0b77dcde-e529-4883-844d-f4344d1b4f56",
    "partnershipId": "f3b9b24d-3d96-4959-88a5-c928931f250f"
}
```
Response Body
```json
{
    "clientApplicationId": "APP1234",
    "partnershipApplicationId": "0b77dcde-e529-4883-844d-f4344d1b4f56",
    "partnershipId": "f3b9b24d-3d96-4959-88a5-c928931f250f",
    "applicationDetails": {
        "status": "NEW_APPLICATION",
        "message": null,
        "lenderLoanId": null,
        "lenderApplicationId": null,
        "dataRequired": {
            "gst": "",
            "banking": ""
        },
        "tasks": [
            {
                "type": "IDENTITY_VERIFICATION",
                "status": "SUCCESS",
                "message": "",
                "data": {
                    "customerDetails": [
                        {
                            "clientCustomerId": "123456789",
                            "customerVerificationStatus": "VERIFIED,",
                            "message": ""
                        }
                    ],
                }
            },
            {
                "type": "DEDUPE",
                "status": "SUCCESS",
                "message": "",
                "data": {}
            },
            {
                "type": "OFFER",
                "status": "APPROVED",
                "message": "",
                "data": {
                }
            },
            {
                "type": "BANK_STATEMENT_VERIFICATION",
                "status": "SUCCESS",
                "message": "",
                "data": {
                    "customerDetails": [
                        {
                            "clientCustomerId": "123456789",
                            "bsaJourneyLink": "https://example.com/bsa_journey",
                            "bsaDetails": [
                                {
                                    "bsaMode": "AA",
                                    "bsaStatus": "SUCCESS",
                                    "message": ""
                                }
                            ]
                        }
                    ]
                }
            },
            {
                "type": "KYC",
                "status": "SUCCESS",
                "message": "",
                "data": {
                    "customerDetails": [
                        {
                            "clientCustomerId": "123456789",
                            "kycDetails": [
                                {
                                    "customerKycStatus": "SUCCESS",
                                    "kycType": "DIGILOCKER",
                                    "failureReason": null,
                                    "nameMatch": false,
                                    "dobMatch": false,
                                    "addressMatch": false,
                                    "selfieMatch": false,
                                    "consentStatus": "CONFIRMED",
                                    "nameAsPerKyc": "Priya Verma",
                                    "dobAsPerKyc": "1995-08-21",
                                    "addressAsPerKyc": "456, Residency Road, Bangalore",
                                    "photoAsPerKyc": "https://example.com/photos/67890.jpg",
                                    "kycId": "DLKYC456789",
                                    "genderAsPerKyc": "FEMALE",
                                    "kycJourneyUrl": "https://example.com/redirect/12345",
                                    "message": null
                                }
                            ]
                        }
                    ]
                }
            },
            {
                "type": "BANK_ACCOUNT_VALIDATION",
                "status": "SUCCESS",
                "message": "",
                "data": {
                    "bankAccountDetails": [
                        {
                            "bankName": "HDFC Bank",
                            "accountHolderName": "Rajesh Kumar Sharma",
                            "bankAccountNumber": "1234567890",
                            "ifscCode": "ABCD0000123",
                            "accountType": "CURRENT",
                            "entityType": "BORROWER",
                            "bankAccountType": [
                                "DISBURSEMENT_ACCOUNT",
                                "REPAYMENT_ACCOUNT"
                            ],
                            "clientCustomerId": "C12345",
                            "clientBankAccountId": "BA1234",
                            "disbursementAmount": 6000,
                            "bankAccountExt": [
                                {
                                    "key": "value"
                                }
                            ],
                            "bankAccountValidationStatus": "SUCCESS",
                            "message": null
                        }
                    ]
                }
            },
            {
                "type": "DOCUMENT_GENERATION",
                "status": "SUCCESS",
                "message": "",
                "data": {}
            },
            {
                "type": "RPS",
                "status": "SUCCESS",
                "message": "",
                "data": {}
            },
            {
                "type": "ESIGNING",
                "status": "SUCCESS",
                "message": "",
                "data": {
                    "documentDetails": [
                        {
                            "documentId": "123e4567-e89b-12d3-a456-426614174000",
                            "esigningDocumentType": "LOANOS_DOCUMENTS",
                            "esigningDocumentSubType": "LOAN_APPLICATION_FORM",
                            "esigningDocumentStatus": "REQUESTED",
                            "message": null,
                            "customerDetails": [
                                {
                                    "clientCustomerId": "CUST1234",
                                    "esigningJourneyUrl": "https://example.com/esigning/CUST1234/DOC1234",
                                    "esigningCustomerStatus": "REQUESTED"
                                }
                            ]
                        }
                    ]
                }
            },
            {
                "type": "DISBURSEMENT",
                "status": "PENDING",
                "data": {
                    "disbursementDetails": [
                        {
                            "entityType": "CO_APPLICANT",
                            "clientCustomerId": "CUST1234",
                            "clientBankAccountId": "BA1234",
                            "bankAccountNumber": "1234567890",
                            "disbursementStatus": "PENDING",
                            "utrNumber": null,
                            "message": null
                        }
                    ]
                }
            },
            {
                "type": "MANDATE",
                "status": "REGISTERED",
                "message": "",
                "data": {
                    "mandateId": "ABCDVMPCV8QTRGGSH8ZW",
                    "umrnId": "UMRN123456",
                    "mandateJourneyUrl": "https://example.com/mandateJourney"
                }
            }
        ]
    }
}
```

### Problem Statement

- An API provides a list of tasks. 
- Basis the provided task's status (PENDING/INPROGRESS/SUCCESS), perform the following:
    - If PENDING skip this task
    - If INPROGRESS/SUCCESS fetch the corresponding data fields from the underlying data system. 
- Write the logic following SOLID principles and any design patterns to make the code extensible following best practices.

Assume the service which provides the list of tasks and the underlying data system as TaskExternalService and DataExternalService respectively.

