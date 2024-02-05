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
      content: string
      contentSections: string[]
      remote: boolean
      createdAt: string
      skills: string[]
      isSaved: boolean
    }
  }

  export type JobDetails = JobDetailsResponse['data']
}
