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

最后更新：2026-03-11 07:50:04 UTC（2026-03-11 15:50:04 UTC+8）

**代理总数：73**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 73 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 45.136.131.47:8443 | ✓ 278ms | ✓ 645ms | ✓ 757ms | ✓ 732ms | ✓ 615ms | http |
| 45.136.131.63:8443 | ✓ 190ms | ✓ 1810ms | ✓ 189ms | ✓ 722ms | ✓ 530ms | http |
| 45.136.130.175:8443 | ✓ 265ms | ✓ 654ms | ✓ 1628ms | ✓ 694ms | ✓ 544ms | http |
| 205.209.118.30:3138 | ✓ 444ms | ✓ 1885ms | ✓ 1140ms | ✓ 1378ms | ✓ 1048ms | http |
| 1.231.81.166:3128 | ✓ 676ms | 否 | ✓ 1141ms | ✓ 956ms | ✓ 786ms | http |
| 86.53.183.16:1080 | ✓ 1107ms | 否 | ✓ 1891ms | 否 | ✓ 1595ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1397ms | 否 | ✓ 1084ms | ✓ 1604ms | http |
| 5.101.0.233:3128 | ✓ 1386ms | 否 | ✓ 1907ms | 否 | ✓ 1823ms | http |
| 120.92.212.16:7890 | ✓ 937ms | 否 | 否 | ✓ 1238ms | ✓ 957ms | http |
| 38.180.2.107:3128 | ✓ 959ms | ✓ 1797ms | ✓ 924ms | ✓ 1858ms | ✓ 1406ms | http |
| 1.225.116.115:1080 | ✓ 1739ms | 否 | ✓ 1022ms | ✓ 1895ms | ✓ 1221ms | http |
| 165.227.5.10:8888 | ✓ 1288ms | 否 | 否 | ✓ 981ms | ✓ 572ms | http |
| 47.77.193.180:1080 | ✓ 741ms | ✓ 1384ms | ✓ 1156ms | 否 | 否 | http |
| 121.138.61.193:8803 | 否 | ✓ 1382ms | ✓ 1514ms | ✓ 1468ms | ✓ 980ms | http |
| 190.9.109.198:999 | ✓ 991ms | ✓ 1474ms | ✓ 1385ms | ✓ 1605ms | ✓ 1375ms | http |
| 158.69.185.37:3129 | ✓ 410ms | 否 | ✓ 1347ms | ✓ 1292ms | ✓ 1508ms | http |
| 45.136.130.191:8443 | ✓ 106ms | ✓ 1251ms | ✓ 392ms | ✓ 668ms | ✓ 504ms | http |
| 45.136.130.188:8443 | ✓ 192ms | ✓ 1141ms | ✓ 421ms | ✓ 665ms | ✓ 504ms | http |
| 35.225.22.61:80 | ✓ 499ms | 否 | ✓ 1108ms | 否 | ✓ 990ms | http |
| 194.213.18.200:443 | ✓ 1720ms | ✓ 1947ms | ✓ 573ms | 否 | 否 | http |
| 101.43.255.96:80 | ✓ 1181ms | ✓ 1274ms | ✓ 939ms | ✓ 1190ms | ✓ 954ms | http |
| 42.96.16.158:1311 | ✓ 1109ms | 否 | ✓ 1152ms | ✓ 1699ms | ✓ 1353ms | http |
| 190.212.131.238:3128 | ✓ 884ms | 否 | 否 | ✓ 1999ms | ✓ 1486ms | http |
| 115.231.181.40:8128 | ✓ 868ms | ✓ 1194ms | ✓ 1234ms | 否 | ✓ 1642ms | http |
| 81.70.169.194:80 | ✓ 1036ms | ✓ 1330ms | ✓ 1978ms | ✓ 1320ms | ✓ 965ms | http |
| 185.191.236.162:3128 | ✓ 1786ms | 否 | ✓ 1700ms | 否 | ✓ 1747ms | http |
| 43.167.227.161:1080 | ✓ 1005ms | 否 | ✓ 1960ms | ✓ 802ms | 否 | http |
| 160.238.65.3:3128 | ✓ 1069ms | ✓ 1637ms | ✓ 1971ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 1152ms | 否 | ✓ 399ms | 否 | ✓ 1690ms | http |
| 202.155.12.161:443 | ✓ 1862ms | 否 | 否 | ✓ 1796ms | ✓ 1144ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1712ms | ✓ 1370ms | ✓ 1305ms | http |
| 160.238.65.8:3128 | ✓ 1366ms | ✓ 1713ms | 否 | ✓ 1622ms | 否 | http |
| 160.238.65.2:3128 | ✓ 1365ms | ✓ 1717ms | 否 | ✓ 1996ms | 否 | http |
| 178.236.245.59:3128 | ✓ 737ms | 否 | ✓ 853ms | 否 | ✓ 1693ms | http |
| 160.238.65.4:3128 | ✓ 1873ms | ✓ 1814ms | 否 | ✓ 1606ms | ✓ 1432ms | http |
| 120.92.212.16:8890 | ✓ 1018ms | 否 | ✓ 1000ms | ✓ 1241ms | ✓ 1236ms | http |
| 178.236.245.17:3128 | ✓ 744ms | ✓ 1757ms | ✓ 739ms | ✓ 1803ms | ✓ 1421ms | http |
| 160.238.65.9:3128 | ✓ 680ms | ✓ 1518ms | ✓ 695ms | 否 | ✓ 1408ms | http |
| 14.225.222.213:7890 | 否 | ✓ 1318ms | 否 | ✓ 1057ms | ✓ 901ms | http |
| 160.238.65.6:3128 | ✓ 674ms | ✓ 1449ms | ✓ 1696ms | 否 | ✓ 1131ms | http |
| 160.238.65.7:3128 | ✓ 1722ms | 否 | ✓ 853ms | ✓ 1566ms | ✓ 1219ms | http |
| 45.136.198.40:3128 | ✓ 1137ms | 否 | ✓ 1469ms | 否 | ✓ 1776ms | http |
| 14.225.222.164:7890 | 否 | 否 | ✓ 1540ms | ✓ 1178ms | ✓ 1229ms | http |
| 160.238.65.5:3128 | ✓ 670ms | 否 | ✓ 1892ms | ✓ 1573ms | ✓ 1222ms | http |
| 152.70.98.46:8888 | ✓ 1458ms | ✓ 1773ms | ✓ 849ms | ✓ 735ms | ✓ 695ms | http |
| 5.252.33.13:2025 | ✓ 1713ms | 否 | ✓ 1389ms | 否 | ✓ 1799ms | http |
| 39.104.201.40:7890 | ✓ 1241ms | ✓ 1810ms | ✓ 987ms | 否 | 否 | http |
| 138.124.90.140:1080 | 否 | 否 | ✓ 1366ms | ✓ 1704ms | ✓ 1484ms | http |
| 172.212.68.37:3128 | ✓ 488ms | ✓ 1621ms | ✓ 955ms | ✓ 1561ms | ✓ 1251ms | http |
| 91.107.141.42:8081 | ✓ 610ms | 否 | ✓ 727ms | 否 | ✓ 1561ms | http |
| 95.3.9.78:3128 | ✓ 1075ms | 否 | ✓ 1250ms | ✓ 1793ms | ✓ 1423ms | http |
| 45.136.130.223:8443 | ✓ 75ms | ✓ 1093ms | ✓ 93ms | ✓ 672ms | ✓ 495ms | http |
| 190.6.54.12:6969 | ✓ 1217ms | ✓ 1712ms | ✓ 1833ms | ✓ 1980ms | ✓ 1527ms | http |
| 95.3.9.78:8080 | ✓ 1032ms | 否 | ✓ 1435ms | 否 | ✓ 1439ms | http |
| 113.177.131.2:3128 | 否 | ✓ 1487ms | ✓ 811ms | ✓ 1055ms | ✓ 860ms | http |
| 114.231.73.23:1080 | ✓ 1270ms | ✓ 1215ms | ✓ 1585ms | 否 | 否 | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1329ms | ✓ 1303ms | ✓ 1271ms | http |
| 162.248.165.72:1080 | ✓ 780ms | 否 | ✓ 1588ms | ✓ 1785ms | 否 | http |
| 180.103.19.233:1080 | ✓ 1300ms | ✓ 1814ms | 否 | ✓ 1604ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1967ms | ✓ 1384ms | ✓ 1722ms | ✓ 1562ms | 否 | http |
| 116.80.82.92:7777 | ✓ 1565ms | 否 | ✓ 1772ms | ✓ 1820ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1719ms | 否 | 否 | ✓ 1728ms | ✓ 1405ms | http |
| 62.234.206.73:3128 | ✓ 1717ms | ✓ 1190ms | 否 | ✓ 1418ms | ✓ 1484ms | http |
| 144.208.127.181:3128 | ✓ 1209ms | 否 | ✓ 1210ms | ✓ 1521ms | 否 | http |
| 47.101.149.27:9010 | ✓ 1738ms | ✓ 1862ms | 否 | ✓ 1502ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1891ms | 否 | ✓ 1271ms | 否 | ✓ 1915ms | http |
| 168.235.110.63:3128 | ✓ 473ms | 否 | ✓ 889ms | ✓ 1466ms | 否 | http |
| 200.174.198.32:8888 | ✓ 1463ms | 否 | ✓ 1771ms | 否 | ✓ 1873ms | http |
| 150.249.255.91:3128 | ✓ 1045ms | 否 | 否 | ✓ 1831ms | ✓ 685ms | http |
| 34.101.184.164:3128 | ✓ 1613ms | 否 | ✓ 1350ms | ✓ 1950ms | ✓ 1191ms | http |
| 45.136.130.239:8443 | 否 | ✓ 1101ms | 否 | ✓ 798ms | ✓ 548ms | http |
| 207.244.244.178:3128 | ✓ 443ms | ✓ 1546ms | ✓ 351ms | ✓ 1242ms | ✓ 857ms | http |
| 116.80.96.101:3172 | ✓ 1673ms | 否 | ✓ 1482ms | 否 | ✓ 1608ms | http |

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
