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

  export interface JobDetailsResponse {
    data: {
      title: string
      salaryFrom: number
      salaryTo: number
      level: string
      company: {
        id: string
        name: string
        image: string
      }
      remote: boolean
      createdAt: string
      skills: string[]
<<<<<<< HEAD
      isSaved: boolean
=======
>>>>>>> 04f36aef4b0e458a5c49573d2629164ff8017be7
    }
  }

  export type JobDetails = JobDetailsResponse['data']
}
