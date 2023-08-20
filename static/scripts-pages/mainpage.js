var originalData = {}; // Объявляем переменную для хранения оригинальных данных

function getOriginalData(cells) {
    const data = {};
    cells.forEach(cell => {
        const columnName = cell.classList[0];
        const originalValue = cell.textContent.trim();
        data[columnName] = originalValue;
    });
    return data;
}

function edit(button) {
    const pensil = button.querySelector('.pensil');
    const row = button.parentNode.parentNode; // Получаем родительскую строку
    const cells = row.querySelectorAll('.name, .addres, .age, .number'); // Получаем ячейки

    const isEditing = pensil.dataset.editing === 'true';

    const allButtons = document.querySelectorAll('.but');
    allButtons.forEach(btn => {
        const btnPensil = btn.querySelector('.pensil');
        if (btn === button) {
            btnPensil.dataset.editing = isEditing ? 'false' : 'true';
            btnPensil.style.backgroundColor = isEditing ? 'yellow' : 'green';

            if (btnPensil.style.backgroundColor === 'green') {
                originalData = getOriginalData(cells);
            }

            btnPensil.src = isEditing ? '../static/images/edit.ico' : '../static/images/check_mark.ico';

            if (isEditing) {
                cells.forEach(cell => {
                    const input = cell.querySelector('input');
                    if (input) {
                        cell.dataset.previousValue = input.value; // Сохраняем предыдущее значение в data атрибуте
                    }
                });
                
                edit_data(cells, originalData); // Передаем cells и originalData в функцию edit_data
                
            } else {
                cells.forEach(cell => {
                    const input = cell.querySelector('input');
                    if (input) {
                        input.value = cell.dataset.previousValue; // Восстанавливаем предыдущее значение из data атрибута
                        delete cell.dataset.previousValue; // Удаляем data атрибут после восстановления
                    }
                });
            }
        } else {
            btnPensil.dataset.editing = 'false';
            btnPensil.style.backgroundColor = 'yellow';
            btnPensil.src = '../static/images/edit.ico';
        }
    });

    if (isEditing) {
        // Применяем изменения
        cells.forEach(cell => {
            const input = cell.querySelector('input');
            if (input) {
                cell.textContent = input.value;
            }
        });
    } else {
        // Сначала вернём ранее редактированные строки в обычное состояние
        const previouslyEditedRows = document.querySelectorAll('.table_data tr.editing');
        previouslyEditedRows.forEach(previouslyEditedRow => {
            const cellsInRow = previouslyEditedRow.querySelectorAll('.name, .addres, .age, .number');
            cellsInRow.forEach(cell => {
                const input = cell.querySelector('input');
                if (input) {
                    cell.textContent = input.value;
                }
            });
            previouslyEditedRow.classList.remove('editing');
        });

        // Превращаем ячейки в поля ввода
        cells.forEach(cell => {
            const input = document.createElement('input');
            input.className = 'input-field';
            input.value = cell.textContent;
            cell.textContent = '';
            cell.appendChild(input);

            input.style.width = (input.value.length + 2) + 'ch';
        });

        row.classList.add('editing'); // Добавляем класс, чтобы помнить, что эта строка в режиме редактирования
    }

    pensil.dataset.editing = isEditing ? 'false' : 'true';
}

function edit_data(cells, data) {
    const rowData = {}; // Создаем объект для данных строки

    cells.forEach(cell => {
        const input = cell.querySelector('input');
        if (input) {
            const columnName = cell.classList[0]; // Получаем класс ячейки, который соответствует названию колонки
            rowData[columnName] = input.value; // Добавляем данные в объект
        }
    });

    console.log(rowData);
    console.log(data);

    // Отправляем POST-запрос
    fetch('/edit', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(rowData)
    })
    .then(response => {
        console.log(); // Выводите объект ответа в консоль
        return response.json();
    })
    .then(data => {
        console.log();
        // Дополнительные действия после получения ответа от сервера
    })
    .catch(error => {
        console.error('Error:', error);
        // Обработка ошибок
    });
}
