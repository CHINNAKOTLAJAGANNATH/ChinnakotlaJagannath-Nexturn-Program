

let expenses = [];
let totalSpending = {};

document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("expense-form");
    form.addEventListener("submit", addExpense);

    const expensesTable = document.getElementById("expenses-table");
    const totalSpendingTable = document.getElementById("total-spending-table");

    loadExpenses();
});

function addExpense(event) {
    event.preventDefault();

    const amount = document.getElementById("amount").value;
    const description = document.getElementById("description").value;
    const category = document.getElementById("category").value;

    const expense = {
        amount: parseFloat(amount),
        description: description,
        category: category
    };

    expenses.push(expense);
    saveExpenses();

    const tableRow = document.createElement("tr");
    tableRow.innerHTML = `
        <td>${expense.amount}</td>
        <td>${expense.description}</td>
        <td>${expense.category}</td>
        <td><button class="delete-btn" data-index="${expenses.length - 1}">Delete</button></td>
    `;
    document.getElementById("expenses-tbody").appendChild(tableRow);

    updateTotalSpending();
    updateTotalSpendingTable();

    document.getElementById("amount").value = "";
    document.getElementById("description").value = "";
    document.getElementById("category").value = "";
}

function loadExpenses() {
    if (localStorage.getItem("expenses")) {
        expenses = JSON.parse(localStorage.getItem("expenses"));
        expenses.forEach((expense, index) => {
            const tableRow = document.createElement("tr");
            tableRow.innerHTML = `
                <td>${expense.amount}</td>
                <td>${expense.description}</td>
                <td>${expense.category}</td>
                <td><button class="delete-btn" data-index="${index}">Delete</button></td>
            `;
            document.getElementById("expenses-tbody").appendChild(tableRow);
        });
        updateTotalSpending();
        updateTotalSpendingTable();
    }
}

function saveExpenses() {
    localStorage.setItem("expenses", JSON.stringify(expenses));
}

function updateTotalSpending() {
    totalSpending = {};
    expenses.forEach(expense => {
        if (!totalSpending[expense.category]) {
            totalSpending[expense.category] = 0;
        }
        totalSpending[expense.category] += expense.amount;
    });
}

function updateTotalSpendingTable() {
    const totalSpendingTableBody = document.getElementById("total-spending-tbody");
    totalSpendingTableBody.innerHTML = "";
    Object.keys(totalSpending).forEach(category => {
        const tableRow = document.createElement("tr");
        tableRow.innerHTML = `
            <td>${category}</td>
            <td>${totalSpending[category]}</td>
        `;
        totalSpendingTableBody.appendChild(tableRow);
    });
}

document.addEventListener("click", function(event) {
    if (event.target.classList.contains("delete-btn")) {
        const index = event.target.dataset.index;
        expenses.splice(index, 1);
        saveExpenses();
        document.getElementById("expenses-tbody").innerHTML = "";
        expenses.forEach((expense, index) => {
            const tableRow = document.createElement("tr");
            tableRow.innerHTML = `
                <td>${expense.amount}</td>
                <td>${expense.description}</td>
                <td>${expense.category}</td>
                <td><button class="delete-btn" data-index="${index}">Delete</button></td>
            `;
            document.getElementById("expenses-tbody").appendChild(tableRow);
        });
        updateTotalSpending();
        updateTotalSpendingTable();
    }
});
