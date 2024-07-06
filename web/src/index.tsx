import { App } from 'app/App';
import { StoreProvider } from 'app/providers/StoreProvider/ui/StoreProvider';
import { ThemeProvider } from 'app/providers/ThemeProvider';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <StoreProvider>
        <BrowserRouter>
            <ThemeProvider>
                <App />
            </ThemeProvider>
        </BrowserRouter>
    </StoreProvider>,
);
