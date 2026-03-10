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

最后更新：2026-03-10 08:37:33 UTC（2026-03-10 16:37:33 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.3.236.202:3128 | ✓ 503ms | 否 | ✓ 1178ms | ✓ 1072ms | ✓ 1063ms | http |
| 120.232.242.119:22222 | 否 | ✓ 1266ms | ✓ 967ms | 否 | ✓ 956ms | http |
| 168.235.110.63:3128 | ✓ 490ms | 否 | ✓ 1288ms | ✓ 1097ms | ✓ 1528ms | http |
| 120.240.29.51:22222 | ✓ 1010ms | 否 | ✓ 1716ms | ✓ 1277ms | ✓ 964ms | http |
| 35.225.22.61:80 | ✓ 815ms | ✓ 1275ms | ✓ 322ms | 否 | 否 | http |
| 205.209.118.30:3138 | ✓ 311ms | ✓ 1251ms | ✓ 1064ms | ✓ 1598ms | ✓ 921ms | http |
| 190.212.131.238:3128 | ✓ 655ms | ✓ 1543ms | ✓ 650ms | ✓ 1606ms | ✓ 1266ms | http |
| 202.155.12.161:443 | ✓ 1734ms | 否 | 否 | ✓ 1758ms | ✓ 1706ms | http |
| 115.231.181.40:8128 | ✓ 1994ms | 否 | ✓ 1125ms | ✓ 1201ms | 否 | http |
| 45.136.130.223:8443 | ✓ 1674ms | 否 | ✓ 1794ms | ✓ 1900ms | ✓ 1452ms | http |
| 183.249.5.117:22222 | ✓ 1202ms | ✓ 1474ms | ✓ 1488ms | 否 | 否 | http |
| 1.231.81.166:3128 | 否 | ✓ 1236ms | ✓ 1017ms | 否 | ✓ 1160ms | http |
| 45.136.131.47:8443 | ✓ 703ms | ✓ 814ms | ✓ 409ms | ✓ 808ms | ✓ 849ms | http |
| 194.213.18.200:443 | ✓ 589ms | ✓ 1928ms | 否 | 否 | ✓ 1232ms | http |
| 46.183.25.8:443 | ✓ 804ms | 否 | 否 | ✓ 1974ms | ✓ 1148ms | http |
| 5.101.0.233:3128 | ✓ 880ms | ✓ 1958ms | ✓ 1775ms | 否 | ✓ 1751ms | http |
| 178.236.245.59:3128 | ✓ 733ms | 否 | ✓ 1264ms | 否 | ✓ 1572ms | http |
| 147.45.251.242:8888 | ✓ 1464ms | 否 | ✓ 1495ms | 否 | ✓ 1947ms | http |
| 190.9.109.198:999 | ✓ 855ms | ✓ 1525ms | ✓ 1416ms | ✓ 1501ms | ✓ 1317ms | http |
| 136.49.34.18:8888 | 否 | 否 | ✓ 292ms | ✓ 1103ms | ✓ 1850ms | http |
| 91.107.141.42:8081 | 否 | 否 | ✓ 1758ms | ✓ 1634ms | ✓ 1553ms | http |
| 37.139.33.145:1080 | ✓ 1170ms | 否 | ✓ 918ms | ✓ 1835ms | ✓ 1306ms | http |
| 101.43.255.96:80 | ✓ 1448ms | ✓ 1312ms | ✓ 1155ms | ✓ 1449ms | ✓ 1435ms | http |
| 81.70.169.194:80 | ✓ 1964ms | ✓ 1459ms | ✓ 1903ms | ✓ 1375ms | 否 | http |
| 116.80.82.229:3172 | ✓ 1585ms | 否 | ✓ 1629ms | 否 | ✓ 1733ms | http |
| 183.249.5.214:22222 | ✓ 856ms | ✓ 1459ms | ✓ 806ms | ✓ 1078ms | ✓ 856ms | http |
| 120.240.35.176:22222 | ✓ 1082ms | 否 | ✓ 1177ms | ✓ 1372ms | ✓ 1947ms | http |
| 129.213.162.27:17777 | 否 | ✓ 1742ms | 否 | ✓ 1461ms | ✓ 1420ms | http |
| 113.59.32.162:22222 | ✓ 1341ms | ✓ 1729ms | 否 | ✓ 1733ms | ✓ 1330ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1257ms | ✓ 1417ms | ✓ 951ms | http |
| 95.3.9.78:3128 | ✓ 1047ms | 否 | ✓ 747ms | ✓ 1699ms | ✓ 1352ms | http |
| 120.238.159.228:22222 | ✓ 1192ms | ✓ 1369ms | ✓ 1057ms | 否 | ✓ 1131ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1309ms | ✓ 1034ms | ✓ 1616ms | ✓ 1049ms | http |
| 120.238.159.229:22222 | 否 | ✓ 1330ms | ✓ 1049ms | ✓ 1204ms | ✓ 991ms | http |
| 117.159.239.58:22222 | ✓ 1092ms | 否 | ✓ 871ms | ✓ 1160ms | 否 | http |
| 120.92.212.16:8890 | 否 | ✓ 1348ms | ✓ 1099ms | 否 | ✓ 1058ms | http |
| 95.3.9.78:8080 | ✓ 941ms | ✓ 1874ms | ✓ 1401ms | ✓ 1742ms | ✓ 1609ms | http |
| 185.191.236.162:3128 | ✓ 1525ms | ✓ 1638ms | ✓ 1436ms | 否 | ✓ 1417ms | http |
| 91.233.223.147:3128 | 否 | 否 | ✓ 849ms | ✓ 1978ms | ✓ 1512ms | http |
| 34.101.184.164:3128 | ✓ 1862ms | 否 | ✓ 1385ms | ✓ 1749ms | ✓ 1210ms | http |
| 104.248.81.109:3128 | ✓ 580ms | ✓ 1742ms | ✓ 1985ms | ✓ 1841ms | ✓ 1402ms | http |
| 103.35.188.243:3128 | ✓ 595ms | 否 | 否 | ✓ 1890ms | ✓ 953ms | http |
| 120.198.141.80:22222 | 否 | 否 | ✓ 1067ms | ✓ 1264ms | ✓ 1014ms | http |
| 101.47.73.135:3128 | ✓ 1813ms | 否 | 否 | ✓ 1352ms | ✓ 1405ms | http |
| 162.248.165.72:1080 | ✓ 1217ms | 否 | ✓ 1962ms | 否 | ✓ 1805ms | http |
| 89.107.10.73:3128 | ✓ 535ms | 否 | ✓ 1175ms | 否 | ✓ 1574ms | http |
| 45.136.198.40:3128 | ✓ 759ms | ✓ 1627ms | ✓ 1440ms | ✓ 1933ms | 否 | http |
| 121.237.181.137:8888 | ✓ 927ms | ✓ 1205ms | 否 | 否 | ✓ 986ms | http |
| 152.70.98.46:8888 | ✓ 710ms | 否 | 否 | ✓ 1053ms | ✓ 931ms | http |
| 150.107.140.238:3128 | ✓ 1710ms | 否 | 否 | ✓ 1232ms | ✓ 1067ms | http |
| 113.59.32.142:22222 | ✓ 1338ms | ✓ 1711ms | ✓ 1497ms | ✓ 1747ms | ✓ 1322ms | http |
| 165.227.5.10:8888 | ✓ 418ms | ✓ 962ms | 否 | ✓ 1140ms | ✓ 1773ms | http |
| 45.140.147.155:1081 | ✓ 1700ms | 否 | ✓ 1176ms | ✓ 1929ms | ✓ 1259ms | http |
| 120.238.159.234:22222 | ✓ 966ms | ✓ 1288ms | ✓ 1077ms | ✓ 1178ms | ✓ 947ms | http |
| 117.159.239.49:22222 | 否 | ✓ 1324ms | ✓ 893ms | ✓ 1188ms | ✓ 951ms | http |
| 45.140.147.155:1082 | ✓ 483ms | ✓ 1589ms | ✓ 1372ms | ✓ 1611ms | ✓ 1383ms | http |
| 45.186.6.104:3128 | ✓ 1939ms | ✓ 1960ms | ✓ 1699ms | 否 | 否 | http |
| 120.240.35.177:22222 | ✓ 1057ms | 否 | ✓ 1770ms | ✓ 1316ms | ✓ 1345ms | http |
| 86.53.183.16:1080 | ✓ 802ms | ✓ 1598ms | ✓ 1638ms | 否 | ✓ 1281ms | http |
| 202.129.206.239:3128 | ✓ 1667ms | 否 | ✓ 1828ms | 否 | ✓ 1508ms | http |
| 138.124.53.25:7443 | ✓ 1709ms | 否 | ✓ 1978ms | 否 | ✓ 1830ms | http |
| 120.240.35.178:22222 | 否 | ✓ 1841ms | ✓ 1151ms | ✓ 1290ms | 否 | http |
| 142.171.157.207:3128 | ✓ 412ms | 否 | ✓ 978ms | ✓ 1333ms | ✓ 1449ms | http |
| 103.166.185.54:3128 | ✓ 1736ms | ✓ 1714ms | ✓ 1345ms | ✓ 1290ms | ✓ 1061ms | http |
| 120.240.29.173:22222 | 否 | 否 | ✓ 978ms | ✓ 1280ms | ✓ 984ms | http |
| 120.240.35.161:22222 | 否 | 否 | ✓ 1209ms | ✓ 1331ms | ✓ 1118ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1384ms | ✓ 990ms | ✓ 1222ms | ✓ 958ms | http |
| 113.59.32.163:22222 | 否 | ✓ 1477ms | ✓ 1197ms | ✓ 1353ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1723ms | 否 | ✓ 1953ms | 否 | ✓ 875ms | http |
| 103.39.51.190:8080 | ✓ 1877ms | 否 | ✓ 1893ms | ✓ 1477ms | ✓ 1502ms | http |
| 178.236.245.17:3128 | ✓ 796ms | ✓ 1776ms | ✓ 1388ms | ✓ 1953ms | 否 | http |
| 94.176.3.43:7443 | ✓ 1358ms | 否 | ✓ 1991ms | 否 | ✓ 1520ms | http |
| 104.129.203.245:10139 | ✓ 910ms | ✓ 1119ms | ✓ 939ms | ✓ 913ms | ✓ 754ms | http |
| 104.129.203.245:10026 | ✓ 1371ms | ✓ 880ms | ✓ 857ms | ✓ 880ms | ✓ 672ms | http |
| 104.129.203.245:10733 | ✓ 900ms | ✓ 960ms | 否 | ✓ 844ms | ✓ 669ms | http |
| 116.80.49.169:3172 | 否 | 否 | ✓ 1606ms | ✓ 1922ms | ✓ 1747ms | http |

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
