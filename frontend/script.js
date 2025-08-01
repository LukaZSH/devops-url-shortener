const form = document.getElementById('shorten-form');
const urlInput = document.getElementById('url-input');
const resultDiv = document.getElementById('result');

form.addEventListener('submit', async (event) => {
    event.preventDefault();

    const longUrl = urlInput.value;
    resultDiv.textContent = 'Encurtando...';
    resultDiv.classList.remove('error');

    try {
        const response = await fetch('/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url: longUrl }),
        });

        if (!response.ok) {
            throw new Error('Falha ao encurtar a URL. Tente novamente.');
        }

        const data = await response.json();
        
        resultDiv.innerHTML = `Sua URL curta é: <a href="${data.short_url}" target="_blank">${data.short_url}</a>`;

    } catch (error) {
        resultDiv.textContent = error.message;
        resultDiv.classList.add('error');
    }
});