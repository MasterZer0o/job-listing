export {}

declare global {
  export interface Offer {
    id: number
    location: string
    remoteAvailable: boolean
    company: {
      name: string
      image: string
    }
    level: string
    title: string
    salaryFrom: string
    salaryTo: string
    currency: string
    postedAt: Date
    skills: string[]
  }
}
