import { Result } from '@chainlink/gauntlet-core'
import { logger, BN, prompt } from '@chainlink/gauntlet-core/dist/utils'
import { SolanaCommand, TransactionResponse } from '@chainlink/gauntlet-solana'
import { PublicKey } from '@solana/web3.js'
import { ORACLES_MAX_LENGTH } from '../../../lib/constants'
import { CONTRACT_LIST, getContract } from '../../../lib/contracts'
import RDD from '../../../lib/rdd'

type Input = {
  oracles: {
    signer: string
    transmitter: string
  }[]
  f: number | string
  proposalId: string
}

export default class ProposeConfig extends SolanaCommand {
  static id = 'ocr2:propose_config'
  static category = CONTRACT_LIST.OCR_2

  static examples = [
    'yarn gauntlet ocr2:propose_config --network=devnet --rdd=[PATH_TO_RDD] --state=EPRYwrb1Dwi8VT5SutS4vYNdF8HqvE7QwvqeCCwHdVLC <PROPOSAL_ID>',
  ]

  makeInput = (userInput): Input => {
    if (userInput) return userInput as Input
    const network = this.flags.network || ''
    const rddPath = this.flags.rdd || ''
    const rdd = RDD.load(network, rddPath)
    const aggregator = RDD.loadAggregator(network, rddPath, this.flags.state)
    const _toHex = (a: string) => Buffer.from(a, 'hex')
    const aggregatorOperators: any[] = aggregator.oracles.map((o) => rdd.operators[o.operator])
    const oracles = aggregatorOperators
      .map((operator) => ({
        transmitter: operator.ocrNodeAddress[0],
        signer: operator.ocr2OnchainPublicKey[0].replace('ocr2on_solana_', ''),
      }))
      .sort((a, b) => Buffer.compare(_toHex(a.signer), _toHex(b.signer)))
    const f = aggregator.config.f
    return {
      oracles,
      f,
      proposalId: this.args[0],
    }
  }

  constructor(flags, args) {
    super(flags, args)
    this.require(!!this.flags.state, 'Please provide flags with "state"')
    this.requireArgs('Please provide a proposalId')
  }

  makeRawTransaction = async (signer: PublicKey) => {
    const ocr2 = getContract(CONTRACT_LIST.OCR_2, '')
    const address = ocr2.programId.toString()
    const program = this.loadProgram(ocr2.idl, address)

    const input = this.makeInput(this.flags.input)

    const proposal = new PublicKey(input.proposalId)

    const oracles = input.oracles.map(({ signer, transmitter }) => ({
      signer: Buffer.from(signer, 'hex'),
      transmitter: new PublicKey(transmitter),
    }))
    const f = new BN(input.f)

    const minOracleLength = f.mul(new BN(3)).toNumber()
    this.require(oracles.length > minOracleLength, `Number of oracles should be higher than ${minOracleLength}`)
    this.require(
      oracles.length <= ORACLES_MAX_LENGTH,
      `Oracles max length is ${ORACLES_MAX_LENGTH}, currently ${oracles.length}`,
    )

    logger.log('Config information:', input)
    const ix = program.instruction.proposeConfig(oracles, f, {
      accounts: {
        proposal,
        authority: signer,
      },
    })

    return [ix]
  }

  execute = async () => {
    const rawTx = await this.makeRawTransaction(this.wallet.publicKey)
    await prompt(`Continue setting config on ${this.args[0].toString()}?`)
    const txhash = await this.signAndSendRawTx(rawTx)
    logger.success(`Config set on tx ${txhash}`)

    return {
      responses: [
        {
          tx: this.wrapResponse(txhash, this.args[0]),
          contract: this.args[0],
        },
      ],
    } as Result<TransactionResponse>
  }
}