import React from 'react'

const DashboadWrapper = ({children}:{ children: React.ReactNode}) => {
  return (
    <div className='flex min-h-screen w-full bg-gray-50 text-gay-900'>
        <main className={'flex w-full flex-col bg-gray-50 dark:bg-dark-900 md:pl-64'}>  
          {/* navbar */}
          {children}
        </main>
    </div>
  )
}

export default DashboadWrapper