namespace libmalino;

/// <summary>
/// An implementation of C's linux/reboot.h.
/// </summary>
#pragma warning disable CS1591
public enum LINUX_REBOOT : uint
{
    MAGIC1 = 0xfee1dead,
    MAGIC2	= 672274793,
    MAGIC2A = 85072278,
    MAGIC2B = 369367448,
    MAGIC2C = 537993216,

    CMD_RESTART = 0x01234567,
    CMD_HALT = 0xCDEF0123,
    CMD_CAD_ON = 0x89ABCDEF,
    CMD_CAD_OFF = 0x00000000,
    CMD_POWER_OFF = 0x4321FEDC,
    CMD_RESTART2 = 0xA1B2C3D4,
    CMD_SW_SUSPEND = 0xD000FCE2,
    CMD_KEXEC = 0x45584543
}
#pragma warning restore CS1591