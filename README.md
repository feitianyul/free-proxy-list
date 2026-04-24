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

最后更新：2026-04-24 17:53:19 UTC（2026-04-25 01:53:19 UTC+8）

**代理总数：50**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | 否 | ✓ 1316ms | ✓ 1363ms | ✓ 1183ms | ✓ 837ms | http |
| 20.27.11.248:8561 | ✓ 1062ms | ✓ 1041ms | ✓ 438ms | ✓ 762ms | ✓ 610ms | http |
| 20.27.15.111:8561 | ✓ 1061ms | ✓ 1070ms | ✓ 463ms | ✓ 730ms | ✓ 591ms | http |
| 20.27.13.35:8561 | ✓ 1125ms | ✓ 870ms | ✓ 512ms | ✓ 864ms | 否 | http |
| 20.27.14.220:8561 | ✓ 1080ms | ✓ 1012ms | ✓ 520ms | ✓ 727ms | ✓ 591ms | http |
| 206.206.126.177:2412 | ✓ 693ms | 否 | ✓ 954ms | ✓ 1006ms | ✓ 745ms | http |
| 1.231.81.166:3128 | ✓ 1383ms | ✓ 1320ms | ✓ 991ms | ✓ 925ms | ✓ 794ms | http |
| 115.231.181.40:8128 | ✓ 1027ms | 否 | ✓ 770ms | ✓ 1950ms | 否 | http |
| 47.85.51.197:1080 | ✓ 348ms | 否 | ✓ 1600ms | ✓ 1227ms | ✓ 1433ms | http |
| 35.225.22.61:80 | 否 | ✓ 1254ms | ✓ 458ms | 否 | ✓ 1293ms | http |
| 207.254.71.62:8088 | ✓ 800ms | ✓ 1979ms | ✓ 940ms | 否 | 否 | http |
| 223.84.151.86:30005 | 否 | 否 | ✓ 1657ms | ✓ 1785ms | ✓ 1548ms | http |
| 218.108.131.186:17890 | ✓ 681ms | ✓ 792ms | ✓ 697ms | ✓ 830ms | ✓ 666ms | http |
| 34.71.229.255:3128 | ✓ 413ms | ✓ 1509ms | ✓ 943ms | 否 | ✓ 1056ms | http |
| 120.92.108.86:7890 | ✓ 1828ms | 否 | ✓ 1893ms | 否 | ✓ 1937ms | http |
| 46.101.95.183:8888 | ✓ 1588ms | ✓ 1976ms | ✓ 1540ms | 否 | 否 | http |
| 203.160.174.83:8091 | ✓ 1449ms | 否 | ✓ 1614ms | 否 | ✓ 1569ms | http |
| 35.180.75.159:11913 | ✓ 990ms | 否 | ✓ 1820ms | 否 | ✓ 1763ms | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 856ms | ✓ 1848ms | ✓ 1817ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1073ms | ✓ 869ms | ✓ 1023ms | ✓ 932ms | http |
| 210.223.44.230:3128 | ✓ 1607ms | 否 | ✓ 1973ms | 否 | ✓ 1125ms | http |
| 121.230.9.198:1080 | ✓ 1534ms | 否 | ✓ 1432ms | ✓ 1247ms | 否 | http |
| 47.105.98.23:3128 | ✓ 679ms | ✓ 1921ms | ✓ 1741ms | 否 | ✓ 1350ms | http |
| 161.35.181.96:999 | ✓ 755ms | ✓ 1273ms | ✓ 617ms | ✓ 1240ms | ✓ 1296ms | http |
| 103.163.103.183:8080 | ✓ 1732ms | 否 | ✓ 1561ms | ✓ 1452ms | ✓ 1405ms | http |
| 103.189.254.28:8080 | ✓ 1728ms | 否 | ✓ 1280ms | ✓ 1471ms | 否 | http |
| 103.248.9.81:8080 | ✓ 1721ms | 否 | 否 | ✓ 1605ms | ✓ 1849ms | http |
| 45.59.122.132:80 | ✓ 1979ms | 否 | ✓ 1035ms | ✓ 1560ms | 否 | http |
| 120.27.224.64:3128 | ✓ 688ms | ✓ 781ms | ✓ 720ms | ✓ 931ms | ✓ 759ms | http |
| 152.42.177.32:8888 | ✓ 935ms | 否 | ✓ 1527ms | ✓ 1242ms | ✓ 1229ms | http |
| 117.236.124.166:3128 | ✓ 1581ms | 否 | ✓ 1660ms | 否 | ✓ 1937ms | http |
| 45.153.231.229:8080 | ✓ 992ms | 否 | ✓ 1265ms | 否 | ✓ 1926ms | http |
| 8.219.97.248:80 | ✓ 1435ms | 否 | ✓ 1086ms | 否 | ✓ 1171ms | http |
| 121.230.8.136:1080 | ✓ 1946ms | ✓ 1127ms | ✓ 966ms | 否 | ✓ 920ms | http |
| 62.113.119.14:8080 | ✓ 834ms | ✓ 1791ms | ✓ 1154ms | ✓ 1747ms | ✓ 1347ms | http |
| 84.47.150.125:1080 | ✓ 848ms | 否 | ✓ 1650ms | 否 | ✓ 1849ms | http |
| 113.176.92.71:3128 | ✓ 1211ms | ✓ 1248ms | ✓ 1201ms | ✓ 1498ms | ✓ 901ms | http |
| 45.140.147.82:1081 | ✓ 642ms | ✓ 1317ms | ✓ 1382ms | 否 | ✓ 1341ms | http |
| 42.101.8.101:8888 | ✓ 1439ms | 否 | ✓ 906ms | ✓ 1424ms | 否 | http |
| 122.224.198.218:808 | ✓ 1927ms | 否 | ✓ 1974ms | 否 | ✓ 1931ms | http |
| 120.92.212.16:8890 | ✓ 803ms | 否 | 否 | ✓ 1922ms | ✓ 1436ms | http |
| 47.84.73.61:1080 | ✓ 1480ms | 否 | ✓ 892ms | ✓ 1070ms | ✓ 809ms | http |
| 202.129.206.239:3128 | ✓ 1356ms | 否 | 否 | ✓ 1982ms | ✓ 1291ms | http |
| 218.60.0.214:80 | 否 | ✓ 1450ms | ✓ 958ms | ✓ 1273ms | ✓ 942ms | http |
| 116.80.62.132:3128 | ✓ 1019ms | 否 | ✓ 1079ms | ✓ 745ms | ✓ 684ms | http |
| 120.92.212.16:7890 | ✓ 1681ms | ✓ 1115ms | ✓ 1368ms | 否 | 否 | http |
| 61.52.131.172:8443 | ✓ 822ms | ✓ 988ms | ✓ 697ms | ✓ 934ms | ✓ 779ms | http |
| 183.232.248.73:7890 | ✓ 850ms | ✓ 1091ms | ✓ 909ms | ✓ 1008ms | ✓ 808ms | http |
| 103.39.51.207:8080 | ✓ 1364ms | 否 | ✓ 1713ms | ✓ 1336ms | ✓ 1333ms | http |
| 152.70.91.193:40000 | ✓ 1942ms | 否 | 否 | ✓ 1710ms | ✓ 1276ms | http |

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
