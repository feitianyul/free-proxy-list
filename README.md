**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-09 21:44:16 UTC（2026-03-10 05:44:16 UTC+8）

**代理总数：74**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 74 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 14.225.212.37:7890 | 否 | ✓ 1290ms | ✓ 1152ms | ✓ 1840ms | ✓ 946ms | http |
| 61.72.221.94:3128 | ✓ 1714ms | 否 | ✓ 1115ms | 否 | ✓ 762ms | http |
| 217.76.245.80:999 | ✓ 991ms | ✓ 1577ms | ✓ 1561ms | ✓ 1709ms | ✓ 1450ms | http |
| 162.240.154.26:3128 | ✓ 590ms | ✓ 1237ms | ✓ 533ms | ✓ 1287ms | ✓ 985ms | http |
| 202.155.12.161:443 | ✓ 1776ms | 否 | 否 | ✓ 1828ms | ✓ 1538ms | http |
| 35.225.22.61:80 | ✓ 858ms | 否 | ✓ 434ms | ✓ 1129ms | ✓ 971ms | http |
| 121.237.181.137:8888 | ✓ 799ms | ✓ 995ms | ✓ 885ms | ✓ 1186ms | ✓ 825ms | http |
| 101.43.255.96:80 | ✓ 889ms | ✓ 1200ms | ✓ 905ms | ✓ 1286ms | ✓ 1034ms | http |
| 115.231.181.40:8128 | ✓ 854ms | ✓ 1106ms | ✓ 868ms | 否 | ✓ 903ms | http |
| 81.70.169.194:80 | ✓ 996ms | ✓ 1237ms | ✓ 1031ms | 否 | ✓ 1006ms | http |
| 61.72.110.114:3128 | ✓ 1193ms | ✓ 1692ms | ✓ 1924ms | ✓ 963ms | ✓ 1707ms | http |
| 1.231.81.166:3128 | ✓ 1688ms | ✓ 1424ms | ✓ 1541ms | ✓ 1428ms | ✓ 1459ms | http |
| 193.168.173.136:443 | ✓ 1104ms | ✓ 1839ms | ✓ 1094ms | ✓ 1887ms | 否 | http |
| 46.183.25.8:443 | ✓ 871ms | 否 | ✓ 210ms | ✓ 807ms | ✓ 665ms | http |
| 8.219.97.248:80 | ✓ 1443ms | 否 | ✓ 1699ms | ✓ 1365ms | ✓ 1586ms | http |
| 152.42.213.210:80 | ✓ 1381ms | 否 | ✓ 940ms | ✓ 1032ms | ✓ 808ms | http |
| 152.42.213.210:8080 | ✓ 1384ms | 否 | ✓ 1937ms | ✓ 1037ms | ✓ 800ms | http |
| 45.186.6.104:3128 | ✓ 1360ms | ✓ 1935ms | ✓ 1686ms | 否 | 否 | http |
| 154.3.236.202:3128 | ✓ 965ms | ✓ 1748ms | ✓ 1330ms | ✓ 1461ms | ✓ 1150ms | http |
| 116.80.96.102:3172 | ✓ 1612ms | 否 | 否 | ✓ 1936ms | ✓ 1769ms | http |
| 114.55.226.123:10086 | ✓ 1682ms | ✓ 1573ms | ✓ 989ms | ✓ 1206ms | ✓ 966ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 845ms | ✓ 1240ms | ✓ 810ms | http |
| 138.124.53.25:7443 | ✓ 1118ms | 否 | 否 | ✓ 1609ms | ✓ 1698ms | http |
| 194.213.18.200:443 | ✓ 945ms | 否 | ✓ 1944ms | 否 | ✓ 1828ms | http |
| 91.233.223.147:3128 | ✓ 1292ms | 否 | ✓ 1991ms | 否 | ✓ 1853ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1240ms | ✓ 1327ms | ✓ 1795ms | http |
| 168.235.110.63:3128 | 否 | 否 | ✓ 1510ms | ✓ 1278ms | ✓ 986ms | http |
| 120.92.212.16:7890 | ✓ 1979ms | 否 | ✓ 1797ms | 否 | ✓ 940ms | http |
| 120.92.212.16:8890 | ✓ 1763ms | 否 | ✓ 959ms | ✓ 1205ms | ✓ 1807ms | http |
| 125.128.12.144:3128 | ✓ 974ms | ✓ 1349ms | ✓ 578ms | ✓ 992ms | ✓ 923ms | http |
| 210.77.29.245:7890 | ✓ 755ms | ✓ 1027ms | ✓ 864ms | ✓ 1011ms | ✓ 786ms | http |
| 136.49.34.18:8888 | 否 | ✓ 1816ms | ✓ 1397ms | 否 | ✓ 1907ms | http |
| 190.9.109.198:999 | ✓ 1170ms | ✓ 1668ms | ✓ 1258ms | ✓ 1770ms | ✓ 1515ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1165ms | ✓ 1206ms | ✓ 1649ms | http |
| 5.101.0.233:3128 | ✓ 870ms | 否 | ✓ 1982ms | 否 | ✓ 1872ms | http |
| 150.249.255.91:3128 | ✓ 1516ms | ✓ 849ms | ✓ 637ms | ✓ 898ms | ✓ 661ms | http |
| 116.80.96.103:3172 | ✓ 1628ms | 否 | 否 | ✓ 1775ms | ✓ 1616ms | http |
| 103.236.89.228:7890 | ✓ 1067ms | ✓ 1340ms | ✓ 1079ms | ✓ 1420ms | ✓ 929ms | http |
| 101.32.244.83:8080 | ✓ 976ms | ✓ 1678ms | ✓ 881ms | ✓ 1125ms | ✓ 1153ms | http |
| 121.43.196.213:8222 | ✓ 928ms | ✓ 1040ms | ✓ 866ms | ✓ 1124ms | ✓ 825ms | http |
| 121.43.196.210:8222 | ✓ 926ms | ✓ 996ms | ✓ 921ms | ✓ 1098ms | ✓ 847ms | http |
| 62.113.119.14:8080 | ✓ 1269ms | 否 | ✓ 1618ms | ✓ 1684ms | ✓ 1358ms | http |
| 116.80.49.167:3172 | ✓ 1474ms | 否 | ✓ 1455ms | ✓ 1835ms | ✓ 1641ms | http |
| 120.238.159.229:22222 | 否 | 否 | ✓ 1248ms | ✓ 1080ms | ✓ 841ms | http |
| 45.136.198.40:3128 | ✓ 837ms | ✓ 1801ms | ✓ 847ms | ✓ 1750ms | ✓ 1340ms | http |
| 117.159.239.49:22222 | ✓ 871ms | ✓ 1019ms | ✓ 788ms | ✓ 1158ms | ✓ 819ms | http |
| 116.80.49.166:3172 | ✓ 1721ms | 否 | 否 | ✓ 1763ms | ✓ 1650ms | http |
| 45.140.147.155:1082 | ✓ 618ms | ✓ 1449ms | ✓ 1579ms | 否 | ✓ 1503ms | http |
| 183.249.5.109:22222 | ✓ 1339ms | ✓ 1275ms | ✓ 949ms | ✓ 1015ms | ✓ 719ms | http |
| 45.140.147.82:1081 | ✓ 723ms | ✓ 1344ms | 否 | ✓ 1762ms | ✓ 1287ms | http |
| 116.80.49.170:3172 | ✓ 1608ms | 否 | ✓ 1832ms | ✓ 1782ms | ✓ 1639ms | http |
| 116.80.49.159:3172 | ✓ 1608ms | 否 | ✓ 1496ms | 否 | ✓ 1658ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1728ms | ✓ 1635ms | ✓ 1588ms | http |
| 183.249.5.111:22222 | 否 | 否 | ✓ 1182ms | ✓ 992ms | ✓ 682ms | http |
| 183.249.5.117:22222 | ✓ 1745ms | 否 | ✓ 624ms | ✓ 968ms | 否 | http |
| 45.140.147.155:1081 | ✓ 1131ms | ✓ 1516ms | ✓ 1228ms | ✓ 1562ms | ✓ 1215ms | http |
| 39.104.201.40:7890 | ✓ 887ms | ✓ 1816ms | 否 | 否 | ✓ 925ms | http |
| 117.159.239.58:22222 | ✓ 827ms | ✓ 976ms | ✓ 782ms | ✓ 1026ms | ✓ 857ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1331ms | ✓ 1023ms | ✓ 798ms | http |
| 61.72.110.54:3128 | ✓ 621ms | 否 | ✓ 832ms | ✓ 974ms | ✓ 863ms | http |
| 101.32.75.4:8888 | ✓ 1502ms | 否 | ✓ 1558ms | ✓ 1057ms | ✓ 953ms | http |
| 121.230.9.26:1080 | ✓ 1286ms | ✓ 1704ms | ✓ 1121ms | ✓ 1618ms | ✓ 1255ms | http |
| 152.70.98.46:8888 | ✓ 1660ms | ✓ 1818ms | ✓ 1942ms | ✓ 1033ms | ✓ 872ms | http |
| 113.132.112.110:9000 | ✓ 1584ms | ✓ 1583ms | ✓ 1422ms | ✓ 1630ms | ✓ 1458ms | http |
| 61.52.131.172:8443 | ✓ 929ms | ✓ 1141ms | ✓ 931ms | ✓ 1158ms | ✓ 886ms | http |
| 103.191.165.146:8090 | 否 | 否 | ✓ 1476ms | ✓ 1333ms | ✓ 1310ms | http |
| 120.238.159.228:22222 | 否 | ✓ 1719ms | ✓ 998ms | ✓ 1081ms | ✓ 855ms | http |
| 113.59.32.162:22222 | ✓ 1017ms | ✓ 1255ms | 否 | ✓ 1894ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1100ms | ✓ 1105ms | 否 | ✓ 1001ms | http |
| 47.77.193.180:1080 | ✓ 644ms | ✓ 642ms | ✓ 31ms | ✓ 660ms | ✓ 502ms | http |
| 117.159.239.52:22222 | ✓ 833ms | ✓ 994ms | ✓ 814ms | ✓ 1131ms | ✓ 889ms | http |
| 120.240.35.178:22222 | ✓ 853ms | ✓ 1197ms | ✓ 1026ms | ✓ 1077ms | ✓ 826ms | http |
| 183.249.5.214:22222 | 否 | 否 | ✓ 635ms | ✓ 1092ms | ✓ 870ms | http |
| 103.217.219.51:1111 | ✓ 1681ms | 否 | ✓ 1436ms | 否 | ✓ 1320ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
