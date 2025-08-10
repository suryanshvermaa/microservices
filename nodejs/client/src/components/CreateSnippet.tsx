import React from 'react'
import { Input } from './ui/input'
import { Textarea } from './ui/textarea'
import { Button } from './ui/button'
import axios, { AxiosError } from "axios";
import { Alert, AlertTitle } from './ui/alert';
import { AlertCircleIcon, CheckCircle2Icon } from 'lucide-react';

interface ISnippet{
  title?: string,
  code?: string,
  description?: string,
  userId?: number
}

interface IAlert{
  alert: boolean,
  success: boolean,
  message: string,
}


const CreateSnippet: React.FC = ({userId}:{userId?: number} ) => {
  const [alert,setAlert]=React.useState<IAlert>({alert:false,success:false,message:""});
  const [snippet,setSnippet]=React.useState<ISnippet|undefined>();
  const createSnippet=async(e:React.FormEvent<HTMLFormElement>)=>{
    try {
      e.preventDefault();
      if(!snippet) return;
      snippet.userId=userId as number;
      const res=await axios.post("http://localhost:3000/api/v1/snippet",{
        ...snippet
      });
      setAlert({
        alert:true,
        success:true,
        message:res.data.message
      });
    } catch (error: AxiosError | any) {
      if(error instanceof AxiosError){
        setAlert({
          alert:true,
          success:false,
          message:error.response?.data.message
        });
        return;
      }
      setAlert({
          alert:true,
          success:false,
          message:error?.message
        });
    }
  }
  return (
    <div>
        <form className='mt-10 flex flex-col gap-4' onSubmit={createSnippet}>
            <Input type='text' placeholder='write a code snippets...' value={snippet?.title} onChange={(e)=>setSnippet({...snippet,title:e.target.value})}/>
            <Textarea placeholder='write a code snippets...' value={snippet?.code} onChange={(e)=>setSnippet({...snippet,code:e.target.value})}/>
            <Textarea placeholder='write a description...' value={snippet?.description} onChange={(e)=>setSnippet({...snippet,description:e.target.value})}/>
            <div className='flex justify-start'>
                <Button type='submit'>Add</Button>
            </div>
            {
              alert.alert&&
                  alert.success&&<Alert>
                      <CheckCircle2Icon />
                      <AlertTitle>{alert.message}</AlertTitle>
                    </Alert>
            }{
                alert.alert && alert.success===false&&<Alert variant={"destructive"}>
                      <AlertCircleIcon />
                      <AlertTitle>{alert.message}</AlertTitle>
                    </Alert>
            }
        </form>
    </div>
  )
}

export default CreateSnippet