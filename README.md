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

最后更新：2026-03-16 12:48:33 UTC（2026-03-16 20:48:33 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 137.220.150.152:6005 | ✓ 1348ms | 否 | ✓ 823ms | ✓ 1595ms | ✓ 899ms | http |
| 202.155.12.161:443 | ✓ 1830ms | 否 | 否 | ✓ 1407ms | ✓ 1135ms | http |
| 1.231.81.166:3128 | ✓ 1836ms | ✓ 1328ms | 否 | ✓ 1091ms | ✓ 834ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 1354ms | ✓ 1689ms | ✓ 1355ms | http |
| 168.235.110.63:3128 | ✓ 374ms | ✓ 1998ms | 否 | ✓ 1339ms | ✓ 853ms | http |
| 45.136.130.209:8448 | ✓ 428ms | ✓ 1548ms | ✓ 1056ms | ✓ 922ms | ✓ 570ms | http |
| 38.34.179.108:8450 | ✓ 430ms | ✓ 1603ms | ✓ 1000ms | ✓ 933ms | ✓ 594ms | http |
| 38.34.179.175:8443 | 否 | 否 | ✓ 138ms | ✓ 995ms | ✓ 673ms | http |
| 45.136.130.238:8447 | 否 | 否 | ✓ 1562ms | ✓ 767ms | ✓ 560ms | http |
| 38.34.179.78:8448 | ✓ 1814ms | 否 | 否 | ✓ 791ms | ✓ 573ms | http |
| 219.117.204.211:7799 | ✓ 1386ms | 否 | ✓ 1247ms | ✓ 977ms | 否 | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1164ms | ✓ 1246ms | ✓ 857ms | http |
| 137.220.151.110:6005 | ✓ 1842ms | 否 | ✓ 1020ms | ✓ 1694ms | ✓ 1085ms | http |
| 38.34.179.60:8450 | ✓ 1605ms | 否 | 否 | ✓ 809ms | ✓ 1289ms | http |
| 14.143.222.113:10155 | ✓ 1740ms | 否 | ✓ 1069ms | ✓ 1468ms | 否 | http |
| 186.148.180.46:999 | ✓ 677ms | 否 | ✓ 662ms | 否 | ✓ 1375ms | http |
| 8.209.239.31:30000 | 否 | ✓ 1080ms | ✓ 608ms | ✓ 754ms | ✓ 692ms | http |
| 124.16.111.161:7890 | ✓ 906ms | ✓ 1207ms | ✓ 1008ms | ✓ 1205ms | ✓ 992ms | http |
| 62.60.177.204:34094 | ✓ 781ms | 否 | 否 | ✓ 1012ms | ✓ 1280ms | http |
| 210.223.44.230:3128 | ✓ 1831ms | ✓ 1058ms | ✓ 973ms | 否 | ✓ 736ms | http |
| 38.145.208.213:8451 | ✓ 674ms | ✓ 1223ms | ✓ 227ms | ✓ 785ms | ✓ 724ms | http |
| 45.136.130.180:8450 | ✓ 669ms | ✓ 1078ms | ✓ 338ms | ✓ 945ms | ✓ 692ms | http |
| 45.136.130.174:8450 | ✓ 671ms | ✓ 924ms | ✓ 652ms | 否 | ✓ 659ms | http |
| 147.161.210.140:8800 | ✓ 1619ms | ✓ 1276ms | ✓ 654ms | ✓ 1416ms | ✓ 1436ms | http |
| 45.136.131.54:8448 | ✓ 672ms | ✓ 1591ms | ✓ 1602ms | ✓ 999ms | 否 | http |
| 101.47.73.135:3128 | ✓ 958ms | 否 | 否 | ✓ 1776ms | ✓ 1131ms | http |
| 38.145.208.163:8450 | ✓ 669ms | ✓ 1701ms | ✓ 186ms | ✓ 743ms | ✓ 1355ms | http |
| 38.34.179.154:8448 | ✓ 669ms | ✓ 1702ms | ✓ 589ms | ✓ 774ms | ✓ 953ms | http |
| 190.242.157.215:8080 | 否 | ✓ 1981ms | ✓ 1491ms | ✓ 1948ms | 否 | http |
| 149.50.116.240:1080 | ✓ 1736ms | ✓ 1870ms | ✓ 1445ms | ✓ 1982ms | ✓ 1509ms | http |
| 113.160.132.26:8080 | ✓ 1013ms | 否 | ✓ 943ms | ✓ 1491ms | ✓ 954ms | http |
| 38.34.179.102:8448 | ✓ 1962ms | 否 | ✓ 1581ms | 否 | ✓ 901ms | http |
| 115.231.181.40:8128 | ✓ 981ms | ✓ 1735ms | ✓ 1005ms | 否 | ✓ 921ms | http |
| 160.250.4.245:1 | ✓ 1299ms | 否 | ✓ 1578ms | ✓ 1225ms | ✓ 1347ms | http |
| 38.34.179.105:8443 | ✓ 1448ms | 否 | ✓ 238ms | ✓ 1065ms | ✓ 1701ms | http |
| 138.124.53.25:7443 | ✓ 1742ms | 否 | ✓ 1951ms | ✓ 1968ms | 否 | http |
| 38.145.203.107:8450 | ✓ 678ms | 否 | ✓ 1128ms | ✓ 947ms | ✓ 680ms | http |
| 120.92.212.16:7890 | ✓ 1328ms | 否 | ✓ 1982ms | 否 | ✓ 1897ms | http |
| 150.107.140.238:3128 | ✓ 1495ms | 否 | ✓ 1992ms | 否 | ✓ 1927ms | http |
| 159.223.42.219:3128 | ✓ 1448ms | 否 | ✓ 741ms | ✓ 1067ms | ✓ 841ms | http |
| 120.92.212.16:8890 | ✓ 1349ms | ✓ 1250ms | 否 | ✓ 1244ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1420ms | 否 | ✓ 1514ms | 否 | ✓ 1515ms | http |
| 101.43.127.100:8877 | ✓ 864ms | ✓ 1320ms | 否 | ✓ 1326ms | ✓ 1910ms | http |
| 35.225.22.61:80 | 否 | ✓ 1974ms | ✓ 893ms | 否 | ✓ 1018ms | http |
| 194.5.212.40:8080 | ✓ 1117ms | 否 | ✓ 1734ms | ✓ 1653ms | ✓ 1465ms | http |
| 106.117.208.101:7890 | ✓ 1165ms | ✓ 1235ms | ✓ 1454ms | ✓ 1364ms | 否 | http |
| 198.24.188.140:30363 | ✓ 815ms | 否 | ✓ 1821ms | ✓ 1699ms | 否 | http |
| 103.113.70.189:1081 | ✓ 1815ms | 否 | 否 | ✓ 1116ms | ✓ 1015ms | http |
| 2.56.122.146:10808 | ✓ 602ms | 否 | ✓ 1471ms | ✓ 1890ms | ✓ 1406ms | http |
| 59.46.216.131:30001 | ✓ 1735ms | ✓ 1390ms | 否 | 否 | ✓ 1962ms | http |
| 116.80.65.81:3172 | ✓ 1601ms | 否 | ✓ 1526ms | ✓ 1842ms | ✓ 1699ms | http |
| 172.212.68.37:3128 | 否 | ✓ 1565ms | ✓ 1409ms | 否 | ✓ 1649ms | http |
| 38.34.178.111:8450 | ✓ 1717ms | ✓ 1907ms | ✓ 819ms | ✓ 990ms | 否 | http |
| 38.145.218.163:8450 | 否 | ✓ 912ms | ✓ 322ms | ✓ 740ms | ✓ 705ms | http |
| 38.34.179.23:8448 | 否 | ✓ 1405ms | ✓ 308ms | ✓ 1781ms | ✓ 569ms | http |
| 38.34.178.244:8450 | ✓ 853ms | ✓ 1022ms | ✓ 149ms | ✓ 748ms | ✓ 558ms | http |
| 38.34.179.14:8450 | ✓ 822ms | 否 | ✓ 539ms | ✓ 1575ms | 否 | http |
| 213.219.214.45:443 | ✓ 792ms | 否 | ✓ 1384ms | 否 | ✓ 1788ms | http |
| 16.78.119.130:443 | ✓ 1900ms | ✓ 1881ms | 否 | 否 | ✓ 1802ms | http |
| 38.34.179.203:8450 | ✓ 706ms | ✓ 1369ms | ✓ 162ms | ✓ 745ms | ✓ 688ms | http |
| 38.145.220.33:8448 | ✓ 849ms | ✓ 1213ms | ✓ 468ms | ✓ 723ms | ✓ 910ms | http |
| 45.136.131.59:8450 | ✓ 1269ms | ✓ 1164ms | ✓ 1148ms | ✓ 1993ms | ✓ 1583ms | http |
| 61.52.131.172:8443 | ✓ 874ms | ✓ 1174ms | ✓ 1037ms | ✓ 1256ms | ✓ 963ms | http |
| 38.145.220.150:8450 | ✓ 368ms | ✓ 959ms | ✓ 384ms | ✓ 924ms | ✓ 627ms | http |
| 47.77.193.180:1080 | ✓ 858ms | ✓ 1826ms | ✓ 497ms | ✓ 768ms | ✓ 566ms | http |

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
