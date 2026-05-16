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

最后更新：2026-05-16 21:44:40 UTC（2026-05-17 05:44:40 UTC+8）

**代理总数：52**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 52 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1791ms | ✓ 1117ms | 否 | ✓ 1460ms | http |
| 218.108.131.186:17890 | ✓ 1015ms | ✓ 1236ms | ✓ 999ms | ✓ 1310ms | ✓ 1062ms | http |
| 45.125.67.37:8443 | ✓ 1166ms | 否 | ✓ 1381ms | ✓ 1431ms | ✓ 1179ms | http |
| 212.58.132.5:8888 | ✓ 1087ms | 否 | ✓ 1538ms | ✓ 1430ms | ✓ 1238ms | http |
| 84.47.150.125:1080 | ✓ 1535ms | 否 | 否 | ✓ 1936ms | ✓ 1775ms | http |
| 5.75.139.30:1081 | ✓ 875ms | ✓ 1724ms | ✓ 558ms | ✓ 1415ms | ✓ 1048ms | http |
| 8.219.97.248:80 | ✓ 1822ms | 否 | ✓ 1870ms | ✓ 1967ms | 否 | http |
| 170.106.136.181:31002 | ✓ 1708ms | ✓ 857ms | ✓ 1739ms | ✓ 957ms | ✓ 1682ms | http |
| 42.114.172.179:2075 | ✓ 1789ms | 否 | ✓ 1712ms | ✓ 1986ms | ✓ 1741ms | http |
| 115.231.181.40:8128 | ✓ 1055ms | 否 | ✓ 1578ms | ✓ 1698ms | ✓ 1872ms | http |
| 103.21.220.141:3128 | ✓ 831ms | 否 | ✓ 839ms | ✓ 1066ms | ✓ 1051ms | http |
| 1.231.81.166:3128 | ✓ 1664ms | ✓ 1581ms | ✓ 1630ms | ✓ 1258ms | ✓ 1027ms | http |
| 148.230.4.241:999 | ✓ 674ms | ✓ 1809ms | ✓ 1360ms | ✓ 1654ms | ✓ 1281ms | http |
| 91.242.229.129:8092 | ✓ 1511ms | 否 | ✓ 1769ms | ✓ 1960ms | ✓ 1979ms | http |
| 185.200.188.234:10001 | ✓ 1811ms | 否 | ✓ 1814ms | 否 | ✓ 1831ms | http |
| 43.156.90.221:10808 | ✓ 1091ms | 否 | 否 | ✓ 1268ms | ✓ 983ms | http |
| 129.212.224.122:3128 | 否 | 否 | ✓ 918ms | ✓ 1273ms | ✓ 1039ms | http |
| 8.154.21.175:3128 | ✓ 1079ms | ✓ 1269ms | ✓ 1074ms | ✓ 1334ms | ✓ 1105ms | http |
| 34.101.184.164:3128 | ✓ 1905ms | 否 | ✓ 1451ms | 否 | ✓ 1372ms | http |
| 57.129.144.178:40000 | ✓ 605ms | ✓ 1617ms | ✓ 1246ms | ✓ 1624ms | ✓ 1485ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1377ms | ✓ 1607ms | ✓ 1314ms | http |
| 137.59.47.73:3128 | ✓ 1610ms | ✓ 1574ms | 否 | ✓ 1663ms | ✓ 1118ms | http |
| 38.180.2.107:3128 | ✓ 785ms | ✓ 1913ms | ✓ 1934ms | 否 | ✓ 1801ms | http |
| 157.0.142.246:10057 | ✓ 1340ms | ✓ 1548ms | ✓ 1188ms | ✓ 1534ms | ✓ 1246ms | http |
| 166.88.55.83:7890 | ✓ 1050ms | ✓ 1344ms | ✓ 827ms | ✓ 1046ms | ✓ 821ms | http |
| 185.121.13.73:3128 | ✓ 1190ms | ✓ 1498ms | ✓ 1710ms | 否 | 否 | http |
| 210.76.192.50:10808 | ✓ 1028ms | ✓ 1402ms | ✓ 1254ms | ✓ 1367ms | ✓ 1073ms | http |
| 185.71.196.92:1080 | ✓ 1215ms | 否 | ✓ 1649ms | 否 | ✓ 1735ms | http |
| 61.155.242.150:5566 | ✓ 1451ms | ✓ 1435ms | ✓ 1468ms | ✓ 1986ms | 否 | http |
| 182.253.109.88:8080 | 否 | 否 | ✓ 1556ms | ✓ 1802ms | ✓ 1754ms | http |
| 158.255.212.55:3256 | ✓ 1893ms | 否 | ✓ 1630ms | ✓ 1706ms | 否 | http |
| 159.89.31.62:8080 | ✓ 475ms | ✓ 1559ms | ✓ 1446ms | 否 | ✓ 1737ms | http |
| 159.223.41.216:9090 | ✓ 918ms | 否 | ✓ 957ms | ✓ 1273ms | ✓ 1020ms | http |
| 47.105.98.23:3128 | ✓ 1073ms | 否 | 否 | ✓ 1434ms | ✓ 1508ms | http |
| 3.101.133.120:80 | 否 | ✓ 1512ms | ✓ 1328ms | ✓ 1283ms | ✓ 1219ms | http |
| 222.127.55.155:8082 | ✓ 1922ms | 否 | 否 | ✓ 1846ms | ✓ 1841ms | http |
| 114.214.170.41:27890 | ✓ 1288ms | ✓ 1635ms | ✓ 1429ms | ✓ 1608ms | ✓ 1303ms | http |
| 160.238.65.4:3128 | 否 | ✓ 1878ms | 否 | ✓ 1610ms | ✓ 1592ms | http |
| 158.255.212.55:9480 | ✓ 1025ms | 否 | ✓ 1324ms | ✓ 1814ms | 否 | http |
| 158.255.212.55:7497 | ✓ 1026ms | 否 | ✓ 1324ms | ✓ 1815ms | 否 | http |
| 158.255.212.55:7839 | ✓ 1018ms | 否 | ✓ 1323ms | ✓ 1813ms | 否 | http |
| 158.255.212.55:9005 | ✓ 1027ms | 否 | ✓ 1322ms | ✓ 1815ms | 否 | http |
| 128.199.116.219:9090 | ✓ 951ms | 否 | ✓ 931ms | ✓ 1658ms | 否 | http |
| 190.14.244.41:9992 | ✓ 1440ms | ✓ 1864ms | ✓ 1611ms | ✓ 1657ms | ✓ 1852ms | http |
| 160.238.65.9:3128 | ✓ 1941ms | 否 | ✓ 819ms | ✓ 1301ms | 否 | http |
| 2.27.32.81:3128 | ✓ 1253ms | ✓ 1789ms | ✓ 995ms | 否 | 否 | http |
| 115.84.169.9:65520 | 否 | 否 | ✓ 1682ms | ✓ 1929ms | ✓ 1943ms | http |
| 129.80.217.21:444 | ✓ 711ms | ✓ 831ms | ✓ 890ms | 否 | ✓ 1700ms | http |
| 121.230.9.26:1080 | ✓ 1280ms | ✓ 1726ms | ✓ 1442ms | ✓ 1934ms | ✓ 1248ms | http |
| 61.52.131.172:8443 | ✓ 1112ms | ✓ 1358ms | ✓ 1176ms | 否 | ✓ 1113ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1911ms | ✓ 874ms | ✓ 1267ms | ✓ 1918ms | http |
| 65.108.203.36:28080 | ✓ 1473ms | 否 | ✓ 1894ms | 否 | ✓ 1915ms | http |

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
