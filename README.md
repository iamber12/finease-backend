
# FinEase Backend

Welcome to the backend repository of FinEase! FinEase is a user-friendly and feature-rich peer-to-peer (P2P) lending and borrowing platform designed to address existing market gaps with secure transactions and efficient loan matching algorithms. Originally hosted on Bitbucket, it has now been imported to Git.

## Overview

This repository contains the backend codebase of FinEase. The backend is responsible for handling various functionalities such as user authentication, loan management, transaction processing, and more. The backend is built using modern development practices and technologies to ensure scalability, security, and maintainability.

Frontend repository: [https://github.com/iamber12/finease-frontend](https://github.com/iamber12/finease-frontend)

## Technologies Used

- **Programming Language:** Go
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Tokens)

## UML Diagram and System Architecture

### Class Diagram

![image](https://github.com/iamber12/finease-p2p-loan-management-app/assets/26606211/8022c5b7-a694-4d83-a9ef-0adab7400f85)

### Architecture Diagram

![image](https://github.com/iamber12/finease-p2p-loan-management-app/assets/26606211/a452bb50-8d4e-451c-99a0-202cfa25eb63)

## Key Components

### 1. User
- **Role:** Represents the users of the system.
- **Relationships:** 
  - Can act as either a Borrower or a Lender.
  - Associated with LoanRequests and LoanProposals.

### 2. Borrower
- **Role:** Represents a user who requests a loan.
- **Relationships:**
  - Linked to the User.
  - Creates LoanRequests.

### 3. Lender
- **Role:** Represents a user who offers a loan.
- **Relationships:**
  - Linked to the User.
  - Creates LoanProposals.

### 4. LoanRequest
- **Role:** Represents a loan request made by a borrower.
- **Relationships:**
  - Linked to the Borrower.
  - Can be linked to multiple LoanProposals.
  - Linked to LoanAgreement.

### 5. LoanProposal
- **Role:** Represents a loan proposal made by a lender.
- **Relationships:**
  - Linked to the Lender.
  - Can be linked to a LoanRequest.
  - Linked to LoanAgreement.

### 6. LoanAgreement
- **Role:** Represents the agreement between a borrower and a lender.
- **Relationships:**
  - Linked to LoanRequest and LoanProposal.
  - Associated with FinancialTransactions.

### 7. FinancialTransaction
- **Role:** Represents the financial transactions involved in the loan process.
- **Relationships:**
  - Linked to LoanAgreement.

For detailed information on endpoints, refer to the [API Controllers](https://github.com/iamber12/finease-backend/tree/dev/pkg/controllers/api).

## Screenshots
![image](https://github.com/iamber12/finease-p2p-loan-management-app/assets/26606211/70307725-8531-4652-a4d6-d3611fe45e80)
