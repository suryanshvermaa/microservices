import React from 'react'
import { Button } from './ui/button'

const Navbar: React.FC = () => {
  return (
    <div className='flex justify-between'>
        <h1 className='font-bold text-2xl'>Code Snippet</h1>
        <Button>Logout</Button>
    </div>
  )
}

export default Navbar