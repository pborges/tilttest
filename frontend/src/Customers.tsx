import { useEffect, useState } from "react"

interface Customer {
    number: number
    name: string
    contactFirstName: string
    contactLastName: string
    phone: string
    addressLine1: string
}

export const DisplayCustomers = () =>{
    const [customers, setCustomers] = useState<Customer[]>([])

    useEffect(() => {
        const loadData = async () => {
            const res = await fetch("http://localhost:8000/customers.json")
            setCustomers(await res.json())
        }
        loadData()
    }, [])

    return <div>
        {customers.map(c=>(<DisplayCustomer key={c.number} customer={c}/>))}
    </div>
}

interface DisplayCustomerProps {
    customer: Customer
}

const DisplayCustomer = ({customer}:DisplayCustomerProps) => {
    return <pre>{JSON.stringify(customer, null, "  ")}</pre>
}