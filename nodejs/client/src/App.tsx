import { AlertDialog, AlertDialogContent } from '@radix-ui/react-alert-dialog'
import './App.css'
import CreateSnippet from './components/CreateSnippet'
import Navbar from './components/Navbar'
import { useEffect, useState } from 'react';
import { Input } from './components/ui/input';
import { Button } from './components/ui/button';
function App() {
    const [userId,setUserId]=useState<number|undefined>();
    const [getUserId,setGetUserId]=useState<boolean>(false);
    useEffect(()=>{
    const userId=localStorage.getItem("userId");
    if(userId){
      setUserId(parseInt(userId));
    }else{
      setGetUserId(true);
    }
  },[])

  const handleUserId=()=>{
    if(!userId) return;
    localStorage.setItem("userId",userId.toString());
    setGetUserId(false);
  }
  return (
      <main>
        <Navbar/>
        <CreateSnippet userId={Number(userId)}/>
        {
          getUserId && <div style={{position:"absolute",top:0,left:0,right:0,bottom:0}} className='h-screen w-screen backdrop-blur-md bg-gray-300/50 dark:bg-gray-800/50 flex justify-center items-center'>
                <AlertDialog defaultOpen>
                  <AlertDialogContent>
                    <h1 className='font-bold text-2xl mb-8'>User Snippet Plateform</h1>
                    <div className='flex flex-col gap-2'>
                      <Input value={userId} placeholder='Enter User Id...' onChange={(e)=>setUserId(parseInt(e.target.value))}/>
                      <Button onClick={handleUserId}>Submit</Button>
                    </div>
                  </AlertDialogContent>
                </AlertDialog>
        </div>
        }
      </main>
  )
}

export default App
