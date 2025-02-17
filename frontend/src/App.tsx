import 'bulma/css/bulma.min.css';
import TopNavBar from './components/TopNavBar';
import MainLayout from './components/MainLayout';

function App() {
  return (
    <div className='main-layout-height'>
      <TopNavBar />
      <MainLayout/>
    </div>
  )
}

export default App;
